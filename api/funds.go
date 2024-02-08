package api

import (
	"merrypay/token"
	"merrypay/types"
	"net/http"
	"strconv"

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

func (s *Server) CompleteWithdrawal(ctx *gin.Context) {
	withdrawalID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorFormat(err.Error()))
		return
	}
	// var req types.CompleteWithdrawalParams
	// if err := ctx.ShouldBindJSON(&req); err != nil {
	// 	ctx.JSON(http.StatusBadRequest, errorFormat(err.Error()))
	// 	return
	// }

	payload := ctx.MustGet(AuthorizationPayloadKey).(*token.Payload)
	if payload.Membership != "admin" {
		ctx.JSON(http.StatusUnauthorized, errorFormat("unauthorized route"))
		return
	}

	err = s.Server.CompleteWithdrawal(ctx, withdrawalID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorFormat(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, "Withdrawal Completed")
}
