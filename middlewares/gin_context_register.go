package middlewares

import (
	"context"

	"github.com/gin-gonic/gin"
)

type ctxKey string

const GIN_CONTEXT_KEY ctxKey = "gin-context"

func GinContextRegister() gin.HandlerFunc {
	return func(ginCtx *gin.Context) {
		// Inject *gin.Context to standard context
		ctx := context.WithValue(ginCtx.Request.Context(), GIN_CONTEXT_KEY, ginCtx)
		ginCtx.Request = ginCtx.Request.WithContext(ctx)

		ginCtx.Next()
	}
}
