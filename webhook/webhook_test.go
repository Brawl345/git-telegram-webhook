package webhook

import (
	"testing"
)

func TestVerifySignature(t *testing.T) {
	t.Run("ValidSignature", func(t *testing.T) {
		payload := []byte("Hello, World!")
		secret := "It's a Secret to Everybody"
		signature := "sha256=757107ea0eb2509fc211221cce984b8a37570b6d7586c22c46f4379c8b043e17"

		if !VerifySignature(payload, secret, signature) {
			t.Errorf("Expected VerifySignature to return true for valid signature")
		}
	})

	t.Run("InvalidSignature", func(t *testing.T) {
		payload := []byte("Goodbye, World!")
		secret := "It's a Secret to Everybody"
		signature := "sha256=757107ea0eb2509fc211221cce984b8a37570b6d7586c22c46f4379c8b043e17"

		if VerifySignature(payload, secret, signature) {
			t.Errorf("Expected VerifySignature to return false for invalid signature")
		}
	})
}
