package api

import (
	"errors"
	"fmt"
	"merrypay/token"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	AuthorizationPayloadKey = "authorization_payload"
	AuthorizationBearerType = "bearer"
)

func authMiddleware(tokenMaker token.Maker) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorization := ctx.GetHeader("Authorization")
		if len(authorization) == 0 {
			err := errors.New("authorization header is not present")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorFormat(err.Error()))
			return
		}
		
		fields := strings.Fields(authorization)
		if len(fields) < 2 {
			err := errors.New("invalid authorization header format")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorFormat(err.Error()))
			return
		}
		
		authorizationType := strings.ToLower(fields[0])
		if authorizationType != AuthorizationBearerType {
			err := fmt.Errorf("unsupported authorization type %s", authorizationType)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorFormat(err.Error()))
			return
		}

		token := fields[1]
		payload, err := tokenMaker.VerifyToken(token)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorFormat(err.Error()))
			return
		}

		ctx.Set(AuthorizationPayloadKey, payload)
		ctx.Next()
	}
}