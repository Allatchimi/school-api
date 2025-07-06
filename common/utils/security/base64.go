package security

import (
	"crypto/rand"
	"encoding/base64"
)

// EncodeBase64 Encodes the input string into Base64 format.
func EncodeBase64(data string) string {
	return base64.StdEncoding.EncodeToString([]byte(data))
}

// DecodeBase64 Decodes a Base64-encoded string and returns an error if the input is invalid.
func DecodeBase64(data string) (string, error) {
	base64Text := make([]byte, base64.StdEncoding.DecodedLen(len(data)))
	n, err := base64.StdEncoding.Decode(base64Text, []byte(data))
	return string(base64Text[:n]), err
}

// GenerateRandomBase64 generates a random base64 URL-safe string of given length.
func GenerateRandomBase64(size int) string {
	bytes := make([]byte, size)
	_, _ = rand.Read(bytes)
	return base64.URLEncoding.EncodeToString(bytes)
}
