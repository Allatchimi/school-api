package utils

import (
	"errors"
	"math/rand"
	"strings"
	"time"
)

const letterAlphaNumericSymbol = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890!@#$%^&*()-=_+"
const letterAlphaNumeric = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
const letterNumeric = "1234567890"

var src = rand.NewSource(time.Now().Unix())

const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

// GenerateRandomCode Generates a random numeric code of specified length.
// Returns the generated code and an error if any.
func GenerateRandomCode(length int) string {
	safeLength := length
	if safeLength <= 0 {
		safeLength = 1
	}
	return GenerateRandomValue(letterNumeric, safeLength)
}

// GenerateRandomPassword Returns a random password of the specified length.
func GenerateRandomPassword(length int) string {
	safeLength := length
	if safeLength <= 0 {
		safeLength = 1
	}
	return GenerateRandomValue(letterAlphaNumericSymbol, safeLength)
}

// GenerateRandomAlphaNumeric Returns a generated alphanumeric
// string with the specified length.
func GenerateRandomAlphaNumeric(length int) string {
	safeLength := length
	if safeLength <= 0 {
		safeLength = 1
	}
	return GenerateRandomValue(letterAlphaNumeric, safeLength)
}

// GenerateRandomValue Returns a random string of specified length, using provided characters.
// It's useful to generate passwords, OTP code and various other things
func GenerateRandomValue(letters string, length int) string {
	sb := strings.Builder{}
	sb.Grow(length)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := length-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letters) {
			sb.WriteByte(letters[idx])
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return sb.String()
}

// GenerateRandomUID returns a unique identifier string.
// It combines a prefix, a random value of a given length, and a suffix.
// Ensures that the generated UID is not in the exclude list.
// Limits the number of attempts to avoid infinite loops.
func GenerateRandomUID(prefix string, acceptedChars string, exclude []string) (string, error) {
	const maxAttempts = 1000

	excludeSet := make(map[string]struct{}, len(exclude))
	for _, uid := range exclude {
		excludeSet[uid] = struct{}{}
	}

	randomLetter := GenerateRandomValue(acceptedChars, 1)

	safeLength := 3
	for range maxAttempts {
		// Generate the random part of the UID
		randomPart := GenerateRandomValue(letterNumeric, safeLength)

		// Build the UID
		var sb strings.Builder
		sb.Grow(len(prefix) + len(randomLetter) + safeLength + 1)
		sb.WriteString(prefix)
		sb.WriteString(randomLetter)
		sb.WriteString(randomPart)
		sb.WriteString(GenerateRandomValue("123456789", 1))

		uid := sb.String()

		// Return if UID is not in exclude list
		if _, exists := excludeSet[uid]; !exists {

			return uid, nil
		}
	}

	return "", errors.New("failed to generate a unique UID after max attempts")
}
