package braintree

import "github.com/lionelbarrow/braintree-go/nullable"

type Customer struct {
	XMLName        string          `xml:"customer"`
	Id             string          `xml:"id,omitempty"`
	FirstName      string          `xml:"first-name,omitempty"`
	LastName       string          `xml:"last-name,omitempty"`
	Company        string          `xml:"company,omitempty"`
	Email          string          `xml:"email,omitempty"`
	Phone          string          `xml:"phone,omitempty"`
	Fax            string          `xml:"fax,omitempty"`
	Website        string          `xml:"website,omitempty"`
	CreditCard     *CreditCard     `xml:"credit-card,omitempty"`
	CreditCards    *CreditCards    `xml:"credit-cards,omitempty"`
	PaypalAccount  *PaypalAccount  `xml:"paypal-account,omitempty"`
	PaypalAccounts *PaypalAccounts `xml:"paypal-accounts,omitempty"`
	ApplePayCard   *ApplePayCard   `xml:"apple-pay-card,omitempty"`
	ApplePayCards  *ApplePayCards  `xml:"apple-pay-cards,omitempty"`
}

// DefaultCreditCard returns the default credit card, or nil
func (c *Customer) DefaultCreditCard() *CreditCard {
	for _, card := range c.CreditCards.CreditCard {
		if card.Default {
			return card
		}
	}
	return nil
}

// NOTE: why not make PaymentMethod an interface instead of an embedded type?
func (c *Customer) PaymentMethods() []*PaymentMethod {
	methods := []*PaymentMethod{}
	if c.CreditCards != nil {
		for _, cc := range c.CreditCards.CreditCard {
			methods = append(methods, &cc.PaymentMethod)
		}
	}
	if c.PaypalAccounts != nil {
		for _, pp := range c.PaypalAccounts.PaypalAccount {
			methods = append(methods, &pp.PaymentMethod)
		}
	}
	if c.ApplePayCards != nil {
		for _, ap := range c.ApplePayCards.ApplePayCard {
			methods = append(methods, &ap.PaymentMethod)
		}
	}
	return methods
}

type CustomerSearchResult struct {
	XMLName           string              `xml:"customers"`
	CurrentPageNumber *nullable.NullInt64 `xml:"current-page-number"`
	PageSize          *nullable.NullInt64 `xml:"page-size"`
	TotalItems        *nullable.NullInt64 `xml:"total-items"`
	Customers         []*Customer         `xml:"customer"`
}
