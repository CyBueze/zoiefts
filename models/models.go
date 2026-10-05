package models

import "time"

type DonationStatus string

const (
	DonationPending DonationStatus = "pending"
	DonationSuccess DonationStatus = "success"
	DonationFailed  DonationStatus = "failed"
)

type Donation struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Reference  string         `gorm:"uniqueIndex;size:100" json:"reference"`
	Name       string         `gorm:"size:150" json:"name"`
	Email      string         `gorm:"size:150;index" json:"email"`
	Phone      string         `gorm:"size:30" json:"phone"`
	AmountKobo int64          `json:"amount_kobo"`
	Currency   string         `gorm:"size:10;default:NGN" json:"currency"`
	Status     DonationStatus `gorm:"size:20;default:pending;index" json:"status"`
	Message    string         `gorm:"size:500" json:"message"`
	Anonymous  bool           `gorm:"default:false" json:"anonymous"`

	PaystackTransactionID string `gorm:"size:100" json:"paystack_transaction_id"`
	PaymentChannel        string `gorm:"size:50" json:"payment_channel"`
}

func (d Donation) AmountNaira() float64 {
	return float64(d.AmountKobo) / 100
}