package securityUtil

import (
	"api/config"
	"runtime"

	"github.com/alexedwards/argon2id"
)

// EncodeArgon2id Applies the Argon2id hashing algorithm to the password and return the resulting hashed string.
func EncodeArgon2id(password string) (string, error) {
	params := &argon2id.Params{
		Memory:      uint32(config.Env.ArgonMemoryLeft * config.Env.ArgonMemoryRight),
		Iterations:  uint32(config.Env.ArgonIterations),
		Parallelism: uint8(runtime.NumCPU()),
		SaltLength:  uint32(config.Env.ArgonSaltLength),
		KeyLength:   uint32(config.Env.ArgonKeyLength),
	}
	tempHash, tempErr := argon2id.CreateHash(password, params)
	if tempErr != nil {
		return "", tempErr
	}
	return EncodeBase64(tempHash), nil
}

// CompareArgon2id Verifies if the Argon2id password matches the string.
func CompareArgon2id(password string, hashedPassword string) (bool, error) {
	initialHashedPassword, _ := DecodeBase64(hashedPassword)
	return argon2id.ComparePasswordAndHash(password, initialHashedPassword)
}
