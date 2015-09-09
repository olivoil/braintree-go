package braintree

import (
	"crypto/hmac"
	"crypto/sha1"
	"errors"
	"fmt"
	"io"
	"strings"
)

type SignatureError struct {
	message string
}

func (i SignatureError) Error() string {
	if i.message == "" {
		return "Invalid Signature"
	}
	return i.message
}

func newHmacer(bt *Braintree) hmacer {
	return hmacer{bt}
}

type hmacer struct {
	*Braintree
}

func (h hmacer) matchingSignaturePair(signatureKeyPair string) (string, string) {
	signaturePairs := strings.Split(signatureKeyPair, "&")
	validPairs := [][]string{}
	for _, pair := range signaturePairs {
		if !strings.Contains(pair, "|") {
			continue
		}
		validPairs = append(validPairs, strings.Split(pair, "|"))
	}
	for _, pair := range validPairs {
		publicKey := pair[0]
		signature := pair[1]
		if publicKey == h.PublicKey {
			return publicKey, signature
		}
	}
	return "", ""
}

func (h hmacer) verifySignature(signature, payload string) (bool, error) {
	signature, err := h.parseSignature(signature)
	if err != nil {
		return false, err
	}
	expectedSignature, err := h.hmac(payload)
	if err != nil {
		return false, err
	}
	return hmac.Equal([]byte(expectedSignature), []byte(signature)), nil
}

func (h hmacer) parseSignature(signatureKeyPair string) (string, error) {
	publicKey, signature := h.matchingSignaturePair(signatureKeyPair)
	if len(publicKey) == 0 {
		return "", SignatureError{"No matching public key in signature"}
	}
	return signature, nil
}

func (h hmacer) hmac(payload string) (string, error) {
	s := sha1.New()
	_, err := io.WriteString(s, h.PrivateKey)
	if err != nil {
		return "", errors.New("Could not write private key to SHA1")
	}
	mac := hmac.New(sha1.New, s.Sum(nil))
	_, err = mac.Write([]byte(payload))
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", mac.Sum(nil)), nil
}
