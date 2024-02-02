package api

import (
	"merrypay/token"
	"merrypay/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) AddBankDetail(ctx *gin.Context) {
	var req types.BankDetailParams
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorFormat(err.Error()))
		return
	}

	payload, ok := ctx.MustGet(AuthorizationPayloadKey).(*token.Payload)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, errorFormat("invalid payload"))
		return
	}
	req.Owner = payload.Username

	err := s.Server.AddBankDetails(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorFormat(err.Error()))
		return
	}

	ctx.JSON(http.StatusCreated, "Bank Detail Added Successfully")
}

func (s *Server) GetBankDetail(ctx *gin.Context) {
	username := ctx.Param("username")

	payload := ctx.MustGet(AuthorizationPayloadKey).(*token.Payload)
	if payload.Username != username && payload.Membership != "admin" {
		ctx.JSON(http.StatusUnauthorized, errorFormat("Unauthorized route"))
		return
	}

	bankDetail, err := s.Server.GetBankDetails(ctx, username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorFormat(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, bankDetail)
}