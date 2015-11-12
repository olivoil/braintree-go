package braintree

import "time"

type ApplePayCard struct {
	PaymentMethod
	ExpirationMonth       string         `xml:"expiration-month,omitempty"`
	ExpirationYear        string         `xml:"expiration-year,omitempty"`
	Expired               bool           `xml:"expired,omitempty"`
	Last4                 string         `xml:"last-4,omitempty"`
	PaymentInstrumentName string         `xml:"payment-instrument-name,omitempty"`
	CardType              string         `xml:"card-type,omitempty"`
	SourceDescription     string         `xml:"source-description,omitempty"`
	Subscriptions         *Subscriptions `xml:"subscriptions,omitempty"`
	CreatedAt             *time.Time     `xml:"created-at,omitempty"`
	UpdatedAt             *time.Time     `xml:"updated-at,omitempty"`
}

type ApplePayCards struct {
	ApplePayCard []*ApplePayCard `xml:"apple-pay-card"`
}
