package braintree

import (
	"testing"
)

func TestCreatePaymentMethodInvalidInput(t *testing.T) {
	cust, err := testGateway.Customer().Create(&Customer{})
	if err != nil {
		t.Fatal(err)
	}

	pm, err := testGateway.PaymentMethod().Create(&PaymentMethod{
		PaymentMethodNonce: "fake-processor-declined-visa-nonce",
		CustomerId:         cust.Id,
	})

	t.Log(pm)

	if err == nil {
		t.Fatal("expected to get error creating payment method because of invalid nonce, but did not")
	}
}

func TestCreatePaymentMethod(t *testing.T) {
	cust, err := testGateway.Customer().Create(&Customer{})
	if err != nil {
		t.Fatal(err)
	}

	pm, err := testGateway.PaymentMethod().Create(&PaymentMethod{
		PaymentMethodNonce: "fake-valid-nonce",
		CustomerId:         cust.Id,
	})

	t.Log(pm)

	// This test should fail because customer id is required
	if err != nil {
		t.Fatal(err)
	}
}
