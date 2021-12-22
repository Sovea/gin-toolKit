package service_hashsalt

import (
	"crypto/rand"
	"crypto/sha512"
	"encoding/base64"
)

// Define salt size
const saltSize = 16

// Generate 16 bytes randomly and securely using the
// Cryptographically secure pseudorandom number generator (CSPRNG)
// in the crypto.rand package
func GenerateRandomSalt(saltSize int) ([]byte, error) {
	var salt = make([]byte, saltSize)

	_, err := rand.Read(salt[:])

	if err != nil {
		return nil, err
	}

	return salt, nil
}

// Combine password and salt then hash them using the SHA-512
// hashing algorithm and then return the hashed password
// as a base64 encoded string
func HashPassword(password string, salt []byte) string {
	// Convert password string to byte slice
	var passwordBytes = []byte(password)

	// Create sha-512 hasher
	var sha512Hasher = sha512.New()

	// Append salt to password
	passwordBytes = append(passwordBytes, salt...)

	// Write password bytes to the hasher
	sha512Hasher.Write(passwordBytes)

	// Get the SHA-512 hashed password
	var hashedPasswordBytes = sha512Hasher.Sum(nil)

	// Convert the hashed password to a base64 encoded string
	var base64EncodedPasswordHash = base64.URLEncoding.EncodeToString(hashedPasswordBytes)

	return base64EncodedPasswordHash
}

// Check if two passwords match
func DoPasswordsMatch(hashedPassword, currPassword string,
	salt []byte) bool {
	var currPasswordHash = HashPassword(currPassword, salt)

	return hashedPassword == currPasswordHash
}
