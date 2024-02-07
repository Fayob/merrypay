package api

import (
	"merrypay/token"
	"merrypay/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) InitiateWithdrawal(ctx *gin.Context) {
	var req types.WithdrawalParam
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorFormat(err.Error()))
		return
	}

	payload := ctx.MustGet(AuthorizationPayloadKey).(*token.Payload)
	req.WithdrawBy = payload.Username
	
	withdrawal, err := s.Server.WithdrawFund(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorFormat(err.Error()))
		return
	}

	ctx.JSON(http.StatusCreated, withdrawal)
}

