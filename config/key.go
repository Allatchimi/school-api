package config

import (
	"api/common/constants"
	"api/common/utils"
)

type Key struct {
	JwtPrivateKey *string
	JwtPublicKey  *string
}

var Keys = &Key{}

// Loads the necessary cryptographic keys.
func LoadKeys() error {
	var errRead error
	// JWT private key
	Keys.JwtPrivateKey, errRead = utils.ReadFileToString(constants.AssetKeysPath + "/jwt/private.pem")
	if errRead != nil {
		return errRead
	}

	// JWT public key
	Keys.JwtPublicKey, errRead = utils.ReadFileToString(constants.AssetKeysPath + "/jwt/public.pem")
	if errRead != nil {
		return errRead
	}
	return errRead
}
