package braintree

import "time"

type PaypalAccount struct {
	PaymentMethod
	Email         string         `xml:"email,omitempty"`
	CreatedAt     *time.Time     `xml:"created-at,omitempty"`
	UpdatedAt     *time.Time     `xml:"updated-at,omitempty"`
	Subscriptions *Subscriptions `xml:"subscriptions,omitempty"`
}

type PaypalAccounts struct {
	PaypalAccount []*PaypalAccount `xml:"paypal-account"`
}

// AllSubscriptions returns all subscriptions for this paypal account, or nil if none present.
func (paypalAccount *PaypalAccount) AllSubscriptions() []*Subscription {
	if paypalAccount.Subscriptions != nil {
		subs := paypalAccount.Subscriptions.Subscription
		if len(subs) > 0 {
			a := make([]*Subscription, 0, len(subs))
			for _, s := range subs {
				a = append(a, s)
			}
			return a
		}
	}
	return nil
}
