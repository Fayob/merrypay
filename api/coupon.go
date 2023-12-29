package api

import (
	"merrypay/token"
	"net/http"

	"github.com/gin-gonic/gin"
)

// type reqParam struct {
// 	Username string `json:"username"`
// }

func (s *Server) GenerateCoupon(ctx *gin.Context) {
	// var req reqParam
	// if err := ctx.ShouldBindJSON(&req); err != nil {
	// 	ctx.JSON(http.StatusBadRequest, errorFormat(err.Error()))
	// 	return
	// }

	payload, ok := ctx.MustGet(AuthorizationPayloadKey).(*token.Payload)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, errorFormat("invalid payload"))
		return
	}

	coupon, err := s.Server.GenerateCoupon(ctx, payload.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorFormat(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, coupon)
}