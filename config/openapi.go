package config

import (
	"api/common/constants"
	"api/common/utils"
)

type OpenAPITemplate struct {
	Redocly   *string
	Scalar    *string
	Stoplight *string
	Swagger   *string
}

var OpenAPITemplates = &OpenAPITemplate{}

// Loads OpenAPI templates from a specified location resources.
func LoadOpenAPITemplates() (err error) {
	// Redocly
	OpenAPITemplates.Redocly, err = utils.ReadFileToString(constants.AssetOpenAPIPath + "/redocly.html")
	if err != nil {
		return
	}

	// Scalar
	OpenAPITemplates.Scalar, err = utils.ReadFileToString(constants.AssetOpenAPIPath + "/scalar.html")
	if err != nil {
		return
	}

	// Stoplight
	OpenAPITemplates.Stoplight, err = utils.ReadFileToString(constants.AssetOpenAPIPath + "/stoplight.html")
	if err != nil {
		return
	}

	// Swagger
	OpenAPITemplates.Swagger, err = utils.ReadFileToString(constants.AssetOpenAPIPath + "/swagger.html")
	if err != nil {
		return
	}

	return
}
