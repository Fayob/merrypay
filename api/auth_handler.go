package api

import (
	"merrypay/types"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

// type UserRequest struct {
// 	Username   string `json:"username" binding:"required"`
// 	Email   string `json:"email" binding:"required"`
// 	FirstName   string `json:"first_name" binding:"required"`
// 	LastName   string `json:"last_name" binding:"required"`
// }

// type authResponse struct {
// 	token string
// 	user  types.UserResponse
// }

func userResponseFunc(user types.User) types.UserResponse {
	return types.UserResponse{
		Username:   user.Username,
		Email:      user.Email,
		FirstName:  user.FirstName,
		LastName:   user.LastName,
		Membership: user.Membership,
		ReferredBy: user.ReferredBy,
		WonJackpot: user.WonJackpot,
		CreatedAt:  user.CreatedAt,
	}
}

func (s *Server) SignUp(ctx *gin.Context) {
	var req types.CreateUserParams
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorFormat(err.Error()))
		return
	}

	user, err := s.Server.RegisterUser(ctx, req)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				ctx.JSON(http.StatusForbidden, errorFormat(err.Error()))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorFormat(err.Error()))
		return
	}

	token, err := s.tokenMaker.CreateToken(user.Username, user.Membership, time.Hour*4)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorFormat(err.Error()))
		return
	}

	response := gin.H{
		"token": token,
		"user":  userResponseFunc(user),
	}

	ctx.JSON(http.StatusCreated, response)
}

type loginParams struct {
	Identifier string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
}

func (s *Server) Login(ctx *gin.Context) {
	var req loginParams
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorFormat(err.Error()))
		return
	}

	user, err := s.Server.LogInUser(ctx, req.Identifier, req.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorFormat(err.Error()))
		return
	}

	token, err := s.tokenMaker.CreateToken(user.Username, user.Membership, time.Hour*4)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorFormat(err.Error()))
		return
	}

	response := gin.H{
		"token": token,
		"user":  userResponseFunc(user),
	}

	ctx.JSON(http.StatusOK, response)
}
