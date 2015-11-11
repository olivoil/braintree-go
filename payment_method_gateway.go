package braintree

type PaymentMethodGateway struct {
	*Braintree
}

func (g *PaymentMethodGateway) Create(pm *PaymentMethod) (*PaymentMethod, error) {
	resp, err := g.execute("POST", "payment_methods", pm)
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode {
	case 201:
		return resp.paymentMethod()
	}
	return nil, &invalidResponseError{resp}
}

func (g *PaymentMethodGateway) Delete(token string) error {
	resp, err := g.execute("DELETE", "payment_methods/any/"+token, nil)
	if err != nil {
		return err
	}
	switch resp.StatusCode {
	case 200:
		return nil
	}
	return &invalidResponseError{resp}
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
