package defens

import (
	"testing"
)

func TestEncryptDecrypt(t *testing.T) {
	testCases := []struct {
		name      string
		key       []byte
		plaintext string
	}{
		{
			name:      "Success",
			key:       []byte("testkey123"),
			plaintext: "Hello, World!",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			encrypted, err := Encrypt(tc.key, tc.plaintext)
			if err != nil {
				t.Fatalf("Encrypt failed: %v", err)
			}

			decrypted, err := Decrypt(tc.key, encrypted)
			if err != nil {
				t.Fatalf("Decrypt failed: %v", err)
			}

			if decrypted != tc.plaintext {
				t.Errorf("Decrypted text does not match original plaintext. Expected: %s, got: %s", tc.plaintext, decrypted)
			}
		})
	}
}
