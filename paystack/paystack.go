package paystack

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const baseURL = "https://api.paystack.co"

type Client struct {
	secretKey string
	http      *http.Client
}

func New(secretKey string) *Client {
	return &Client{
		secretKey: secretKey,
		http:      &http.Client{Timeout: 15 * time.Second},
	}
}

type InitializeRequest struct {
	Email       string `json:"email"`
	AmountKobo  int64  `json:"amount"`
	Reference   string `json:"reference"`
	CallbackURL string `json:"callback_url"`
	Currency    string `json:"currency,omitempty"`
	Metadata    any    `json:"metadata,omitempty"`
}

type InitializeResponseData struct {
	AuthorizationURL string `json:"authorization_url"`
	AccessCode       string `json:"access_code"`
	Reference        string `json:"reference"`
}

type apiResponse[T any] struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

func (c *Client) Initialize(req InitializeRequest) (*InitializeResponseData, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequest(http.MethodPost, baseURL+"/transaction/initialize", bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	c.setHeaders(httpReq)

	resp, err := c.http.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var parsed apiResponse[InitializeResponseData]
	if err := json.NewDecoder(resp.Body).Decode(&parsed); err != nil {
		return nil, err
	}
	if !parsed.Status {
		return nil, fmt.Errorf("paystack initialize failed: %s", parsed.Message)
	}
	return &parsed.Data, nil
}

type VerifyResponseData struct {
	Status    string `json:"status"`
	Reference string `json:"reference"`
	Amount    int64  `json:"amount"`
	Currency  string `json:"currency"`
	Channel   string `json:"channel"`
	ID        int64  `json:"id"`
	Customer  struct {
		Email string `json:"email"`
	} `json:"customer"`
}

func (c *Client) Verify(reference string) (*VerifyResponseData, error) {
	httpReq, err := http.NewRequest(http.MethodGet, baseURL+"/transaction/verify/"+reference, nil)
	if err != nil {
		return nil, err
	}
	c.setHeaders(httpReq)

	resp, err := c.http.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var parsed apiResponse[VerifyResponseData]
	if err := json.NewDecoder(resp.Body).Decode(&parsed); err != nil {
		return nil, err
	}
	if !parsed.Status {
		return nil, fmt.Errorf("paystack verify failed: %s", parsed.Message)
	}
	return &parsed.Data, nil
}

func (c *Client) ValidateWebhookSignature(rawBody []byte, signatureHeader string) bool {
	mac := hmac.New(sha512.New, []byte(c.secretKey))
	mac.Write(rawBody)
	expected := hex.EncodeToString(mac.Sum(nil))
	return hmac.Equal([]byte(expected), []byte(signatureHeader))
}

func (c *Client) setHeaders(req *http.Request) {
	req.Header.Set("Authorization", "Bearer "+c.secretKey)
	req.Header.Set("Content-Type", "application/json")
}

func ReadAndRestoreBody(r *http.Request) ([]byte, error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	r.Body = io.NopCloser(bytes.NewReader(body))
	return body, nil
}