package app

import (
	"testing"
)

func TestEncryptDecrypt(t *testing.T) {
	plaintext := "8f8668b398a84296a26d6a2f294b344c"
	encrypted, err := Crypto().Encrypt(plaintext)
	if err != nil {
		t.Errorf("Error occurred [%v]", err)
	}
	decrypted, err := Crypto().Decrypt(encrypted)
	if err != nil {
		t.Errorf("Error occurred [%v]", err)
	}
	if decrypted != plaintext {
		t.Errorf("Expected decrypted [%v], got [%v]", plaintext, decrypted)
	}
}
