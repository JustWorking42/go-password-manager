package defens

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
			assert.NoError(t, err, "Encrypt should not return an error")

			decrypted, err := Decrypt(tc.key, encrypted)
			assert.NoError(t, err, "Decrypt should not return an error")

			assert.Equal(t, tc.plaintext, decrypted, "Decrypted text does not match original plaintext")
		})
	}
}
