package braintree

import (
	"encoding/base64"
	"encoding/xml"
	"errors"
	"regexp"
)

type WebhookNotificationGateway struct {
	*Braintree
}

func (w *WebhookNotificationGateway) Parse(signature, payload string) (*WebhookNotification, error) {
	hmacer := newHmacer(w.Braintree)
	if verified, err := hmacer.verifySignature(signature, payload); err != nil {
		return nil, err
	} else if !verified {
		return nil, SignatureError{}
	}

	xmlNotification, err := base64.StdEncoding.DecodeString(payload)
	if err != nil {
		return nil, err
	}

	var n WebhookNotification
	err = xml.Unmarshal(xmlNotification, &n)
	if err != nil {
		return nil, err
	}
	return &n, nil
}

func (w *WebhookNotificationGateway) Verify(challenge string) (string, error) {
	matched, _ := regexp.MatchString("^[a-f0-9]{20,32}$", challenge)
	if !matched {
		return "", errors.New("Challenge contains non-hex characters")
	}
	h := newHmacer(w.Braintree)
	digest, err := h.hmac(challenge)
	if err != nil {
		return "", err
	}
	return w.Braintree.PublicKey + "|" + digest, nil
}
