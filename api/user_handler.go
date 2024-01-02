package api

import (
	"merrypay/token"
	"merrypay/types"
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

func (s *Server) UpdateUserProfile(ctx *gin.Context) {
	var req types.UpdateUserParams
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorFormat(err.Error()))
		return
	}

	payload := ctx.MustGet(AuthorizationPayloadKey).(*token.Payload)
	if payload.Username != req.Username && payload.Membership != "admin" {
		ctx.JSON(http.StatusUnauthorized, errorFormat("unauthorized route"))
		return
	}

	err := s.Server.UpdateUser(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorFormat(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, "User Updated Successfully")
}
