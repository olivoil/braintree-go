package braintree

type PaymentMethod struct {
	CustomerId         string             `xml:"customer-id,omitempty"`
	Token              string             `xml:"token,omitempty"`
	PaymentMethodNonce string             `xml:"payment-method-nonce,omitempty"`
	ImageURL           string             `xml:"image-url,omitempty"`
	Default            bool               `xml:"default,omitempty"`
	Options            *CreditCardOptions `xml:options,omitempty"`
}
