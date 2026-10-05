package handlers

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"zoie/config"
	"zoie/models"
	"zoie/paystack"
	"zoie/views/components"
	"zoie/views/pages"
)

type DonateHandler struct {
	DB       *gorm.DB
	Cfg      *config.Config
	Paystack *paystack.Client
}

func NewDonateHandler(db *gorm.DB, cfg *config.Config) *DonateHandler {
	return &DonateHandler{
		DB:       db,
		Cfg:      cfg,
		Paystack: paystack.New(cfg.PaystackSecretKey),
	}
}

type donateForm struct {
	Amount    float64 `form:"amount"`
	Name      string  `form:"name"`
	Email     string  `form:"email"`
	Phone     string  `form:"phone"`
	Message   string  `form:"message"`
	Anonymous bool    `form:"anonymous"`
}

// Initiate handles the htmx form POST from the donation form. It
// validates input, creates a pending Donation row, starts a Paystack
// transaction, and either sends an HX-Redirect header (htmx does a
// full browser redirect) or swaps an error message into #donate-error.
func (h *DonateHandler) Initiate(c *gin.Context) {
	fail := func(message string) {
		render(c, http.StatusOK, components.DonateError(message))
	}

	var in donateForm
	if err := c.ShouldBind(&in); err != nil {
		fail("We couldn't read your donation details. Please try again.")
		return
	}

	in.Name = strings.TrimSpace(in.Name)
	in.Email = strings.TrimSpace(in.Email)

	if in.Amount < 100 {
		fail("Please enter an amount of at least ₦100.")
		return
	}
	if in.Name == "" || in.Email == "" {
		fail("Please fill in your name and email.")
		return
	}

	reference, err := generateReference()
	if err != nil {
		fail("Something went wrong on our end. Please try again.")
		return
	}

	amountKobo := int64(in.Amount * 100)

	donation := models.Donation{
		Reference:  reference,
		Name:       in.Name,
		Email:      in.Email,
		Phone:      in.Phone,
		AmountKobo: amountKobo,
		Currency:   "NGN",
		Status:     models.DonationPending,
		Message:    in.Message,
		Anonymous:  in.Anonymous,
	}
	if err := h.DB.Create(&donation).Error; err != nil {
		log.Printf("failed to create donation record: %v", err)
		fail("Something went wrong on our end. Please try again.")
		return
	}

	initResp, err := h.Paystack.Initialize(paystack.InitializeRequest{
		Email:       in.Email,
		AmountKobo:  amountKobo,
		Reference:   reference,
		CallbackURL: h.Cfg.BaseURL + "/donate/callback",
		Currency:    "NGN",
		Metadata: map[string]any{
			"donor_name": in.Name,
			"anonymous":  in.Anonymous,
		},
	})
	if err != nil {
		log.Printf("paystack initialize failed for %s: %v", reference, err)
		fail("We couldn't reach the payment processor. Please try again shortly.")
		return
	}

	// htmx intercepts this header and does a full-page redirect to Paystack checkout.
	c.Header("HX-Redirect", initResp.AuthorizationURL)
	c.Status(http.StatusOK)
}

func (h *DonateHandler) Callback(c *gin.Context) {
	reference := c.Query("reference")
	if reference == "" {
		reference = c.Query("trxref")
	}

	var donation models.Donation
	if err := h.DB.Where("reference = ?", reference).First(&donation).Error; err != nil {
		c.String(http.StatusNotFound, "Donation not found")
		return
	}

	if donation.Status == models.DonationPending {
		h.verifyAndUpdate(&donation)
	}

	render(c, http.StatusOK, pages.DonateResult(donation))
}

func (h *DonateHandler) Webhook(c *gin.Context) {
	body, err := paystack.ReadAndRestoreBody(c.Request)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	signature := c.GetHeader("x-paystack-signature")
	if !h.Paystack.ValidateWebhookSignature(body, signature) {
		c.Status(http.StatusUnauthorized)
		return
	}

	var event struct {
		Event string `json:"event"`
		Data  struct {
			Reference string `json:"reference"`
		} `json:"data"`
	}
	if err := c.ShouldBindJSON(&event); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	if event.Event == "charge.success" && event.Data.Reference != "" {
		var donation models.Donation
		if err := h.DB.Where("reference = ?", event.Data.Reference).First(&donation).Error; err == nil {
			h.verifyAndUpdate(&donation)
		}
	}

	c.Status(http.StatusOK)
}

func (h *DonateHandler) verifyAndUpdate(donation *models.Donation) {
	result, err := h.Paystack.Verify(donation.Reference)
	if err != nil {
		log.Printf("verify failed for %s: %v", donation.Reference, err)
		return
	}

	if result.Status == "success" {
		donation.Status = models.DonationSuccess
	} else {
		donation.Status = models.DonationFailed
	}
	donation.PaystackTransactionID = fmt.Sprintf("%d", result.ID)
	donation.PaymentChannel = result.Channel

	if err := h.DB.Save(donation).Error; err != nil {
		log.Printf("failed to save donation update for %s: %v", donation.Reference, err)
	}
}

func generateReference() (string, error) {
	b := make([]byte, 8)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return fmt.Sprintf("ZF-%d-%s", time.Now().Unix(), hex.EncodeToString(b)), nil
}