package httpHelper

import (
	"fmt"

	"api/common/types"
	securityUtil "api/common/utils/security"

	"github.com/gin-gonic/gin"
)

// GetJwtContextFromQueryGin Extracts JWT token from query
func GetJwtContextFromQueryGin(ctx *gin.Context) (*types.JwtToken, error) {
	token := ctx.Query("token")
	if len(token) < 1 {
		errMsg := "No token found! Please enter valid information."
		return nil, fmt.Errorf("%s", errMsg)
	}
	jwtToken, _, err := securityUtil.ValidateAuthToken(token)
	if err != nil || jwtToken == nil || jwtToken.UserID < 1 {
		errMsg := "Invalid token! Please enter valid information."
		return nil, fmt.Errorf("%s", errMsg)
	}
	return jwtToken, nil
}
