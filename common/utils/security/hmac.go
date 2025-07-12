package securityUtil

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/base64"
	"encoding/hex"
	"errors"
)

// GenerateHMAC_SHA256_Hex Generates an HMAC-SHA256 signature for the given message using the provided private key.
// Returns the hex-encoded signature as a string and any error encountered.
func GenerateHMAC_SHA256_Hex(message string, privateKey string) (string, error) {
	mac := hmac.New(sha256.New, []byte(privateKey))
	_, err := mac.Write([]byte(message))
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(mac.Sum(nil)), nil
}

// GenerateHMAC_SHA256_Base64URL Generates an HMAC-SHA256 signature for the given message using the provided private key.
// Returns the base64url-encoded signature as a string and any error encountered.
func GenerateHMAC_SHA256_Base64URL(message, privateKey string) (string, error) {
	hexSig, err := GenerateHMAC_SHA256_Hex(message, privateKey)
	if err != nil {
		return "", err
	}

	rawSig, err := hex.DecodeString(hexSig)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(rawSig), nil
}

// VerifyHMAC_SHA256_Hex verifies that the provided HMAC-SHA256 token is valid.
// It takes the provided token (hex-encoded), and the private key used for verification.
// Returns true if the token is valid, false otherwise.
func VerifyHMAC_SHA256_Hex(message, privateKey string, providedToken string) (bool, error) {
	// Recalculate the expected HMAC based on the message and the private key
	expectedHMAC, err := GenerateHMAC_SHA256_Hex(message, privateKey)
	if err != nil {
		return false, err
	}

	// Decode both provided and expected tokens from hex to raw bytes
	providedBytes, err := hex.DecodeString(providedToken)
	if err != nil {
		return false, errors.New("invalid provided token format")
	}

	expectedBytes, err := hex.DecodeString(expectedHMAC)
	if err != nil {
		return false, errors.New("failed to decode expected token")
	}

	// Constant-time comparison to prevent timing attacks
	if subtle.ConstantTimeCompare(providedBytes, expectedBytes) == 1 {
		return true, nil
	}
	return false, nil
}

// VerifyHMAC_SHA256_Base64URL verifies that the provided Base64 URL Safe HMAC-SHA256 token is valid.
// It takes the original message, the provided token, and the private key used for verification.
// Returns true if the token is valid, false otherwise.
func VerifyHMAC_SHA256_Base64URL(message, providedToken, privateKey string) (bool, error) {
	// Recalculate the expected HMAC based on the message and private key
	expectedToken, err := GenerateHMAC_SHA256_Base64URL(message, privateKey)
	if err != nil {
		return false, err
	}

	// Decode both tokens from Base64 URL Safe to raw bytes
	providedBytes, err := base64.URLEncoding.DecodeString(providedToken)
	if err != nil {
		return false, err
	}

	expectedBytes, err := base64.URLEncoding.DecodeString(expectedToken)
	if err != nil {
		return false, err
	}

	// Constant-time comparison to prevent timing attacks
	if subtle.ConstantTimeCompare(providedBytes, expectedBytes) == 1 {
		return true, nil
	}
	return false, nil
}
