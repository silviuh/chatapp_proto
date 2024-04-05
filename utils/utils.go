package utils

import "golang.org/x/crypto/bcrypt"

// HashPassword generates a bcrypt hash of the password.
func HashPassword(password string) (string, error) {
	// The cost of processing the hashing of the password is 14, which is a good balance between security and performance.
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// CheckPasswordHash compares a bcrypt hashed password with its possible plaintext equivalent.
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
