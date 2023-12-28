package api

import (
	"math"
	"merrypay/token"
	"merrypay/types"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

type DashBoard struct {
	User            types.UserResponse
	Earning         types.Earning
	ReferralHistory []types.RefHisResponse
}

type DashBoardParams struct {
	Username string
	Token    string
}

func (s *Server) DashBoard(ctx *gin.Context) {
	// var username reqParam
	// if err := ctx.ShouldBindJSON(&username); err != nil {
	// 	ctx.JSON(http.StatusBadRequest, errorFormat(err.Error()))
	// 	return
	// }
	payload := ctx.MustGet(AuthorizationPayloadKey).(*token.Payload)
	var dashBoard DashBoard
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		user, err := s.Server.GetUser(ctx, payload.Username)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorFormat(err.Error()))
			// return
		}
		defer wg.Done()
		dashBoard.User = userResponseFunc(user)
	}()

	wg.Add(1)
	go func() {
		earning, err := s.Server.GetEarning(ctx, payload.Username)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorFormat(err.Error()))
			// return
		}
		defer wg.Done()
		earning.MediaTotalWithdrawal = int(math.Abs(float64(earning.MediaTotalWithdrawal)))
		earning.ReferralTotalWithdrawal = int(math.Abs(float64(earning.ReferralTotalWithdrawal)))
		dashBoard.Earning = earning
	}()

	wg.Add(1)
	go func() {
		refHistory, err := s.Server.UserReferred(ctx, payload.Username)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorFormat(err.Error()))
			// return
		}
		defer wg.Done()
		dashBoard.ReferralHistory = refHistory
	}()
	wg.Wait()

	ctx.JSON(http.StatusOK, dashBoard)
}
