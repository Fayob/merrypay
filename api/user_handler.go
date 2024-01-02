package api

import (
	"merrypay/token"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) GetUserProfile(ctx *gin.Context) {
	username := ctx.Param("username")
	if username == "" {
		ctx.JSON(http.StatusBadRequest, errorFormat("username cannot be empty"))
		return
	}

	payload := ctx.MustGet(AuthorizationPayloadKey).(*token.Payload)
	if payload.Username != username && payload.Membership != "admin" {
		ctx.JSON(http.StatusUnauthorized, errorFormat("unauthorized route"))
		return
	}

	user, err := s.Server.GetUser(ctx, username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorFormat(err.Error()))
		return 
	}

	ctx.JSON(http.StatusOK, user)
}
