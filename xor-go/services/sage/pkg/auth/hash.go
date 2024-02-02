package hash

import (
	"crypto/rand"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
)

func CreateHash(password string, salt []byte) string {
	passwordBytes := append([]byte(password), salt...)

	sha512Hasher := sha512.New()
	sha512Hasher.Write(passwordBytes)
	hashedPasswordBytes := sha512Hasher.Sum(nil)

	return hex.EncodeToString(hashedPasswordBytes)
}

func CreateSalt(saltSize int) ([]byte, error) {
	salt := make([]byte, saltSize)
	_, err := rand.Read(salt[:])
	if err != nil {
		return nil, fmt.Errorf("failed to create salt: %v", err)
	}

	return salt, nil
}
