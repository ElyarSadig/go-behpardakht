package payment

import (
	"testing"
)

func TestPrepareSOAPRequest(t *testing.T) {
	userId := "testUser"
	password := "testPassword"
	pr, err := NewPaymentRequest("12345", "payer123", "https://callbackurl.com", 1000)
	if err != nil {
		t.Fatalf("NewPaymentRequest returned an error: %v", err)
	}
	payload, err := pr.PrepareSOAPRequest(userId, password)
	if err != nil {
		t.Fatalf("PrepareSOAPRequest returned an error: %v", err)
	}
	if len(payload) == 0 {
		t.Fatalf("Expected non-empty payload, got empty")
	}
}
