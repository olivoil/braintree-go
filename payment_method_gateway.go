package braintree

type PaymentMethodGateway struct {
	*Braintree
}

func (g *PaymentMethodGateway) Find(token string) (*PaymentMethod, error) {
	resp, err := g.execute("GET", "payment_methods/any/"+token, nil)
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode {
	case 200:
		return resp.paymentMethod()
	}
	return nil, &invalidResponseError{resp}
}
