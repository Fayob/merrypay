package api

import (
	"fmt"
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

func (s *Server) Jackpot(ctx *gin.Context) {
	var req types.JackpotParam
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorFormat(err.Error()))
		return
	}

	payload := ctx.MustGet(AuthorizationPayloadKey).(*token.Payload)
	req.Username = payload.Username

	result, err := s.Server.Jackpot(ctx, req.Guess, req.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorFormat(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (s *Server) GetWithdrawalByID(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		fmt.Println("unable to convert params id to int")
		return
	}

	withdrawal, err := s.Server.GetWithdrawReceiptByID(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorFormat(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, withdrawal)
}

func (s *Server) GetWithdrawalsByStatus(ctx *gin.Context) {
	payload := ctx.MustGet(AuthorizationPayloadKey).(*token.Payload)
	if payload.Membership != "admin" {
		ctx.JSON(http.StatusUnauthorized, "unauthorized route")
		return
	}
	status := ctx.Param("status")

	withdrawal, err := s.Server.GetWithdrawalByStatus(ctx, status)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorFormat(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, withdrawal)
}