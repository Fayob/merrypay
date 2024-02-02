package api

import (
	"merrypay/middleware"
	"merrypay/port"
	"merrypay/token"
	"os"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Server port.Server
	tokenMaker token.Maker
	router *gin.Engine
}

func NewHandler(serverPort port.Server) (*Server, error) {
	symmetricKey := os.Getenv("SYMMETRIC_KEY")
	tokenMaker, err := middleware.NewPasetoMaker(symmetricKey)
	if err != nil {
		return nil, err
	}
	server := &Server{
		Server: serverPort,
		tokenMaker: tokenMaker,
	}

	server.setupRoute()

	return server, nil
}

func (s *Server) setupRoute() {
	router := gin.Default()

	// Auth
	router.POST("/register", s.SignUp)
	router.POST("/login", s.Login)

	app := router.Group("/").Use(authMiddleware(s.tokenMaker))
	
	// Generate Coupon by admin
	app.GET("/coupon", s.GenerateCoupon)

	// User
	app.GET("/user/dashboard", s.DashBoard)
	app.GET("/user/:username", s.GetUserProfile)
	app.PATCH("/user/update", s.UpdateUserProfile)
	app.PATCH("/user/membership", s.UpdateUserMemberShip)
	app.PATCH("/updatePassword", s.UpdatePassword)

	// Bank Details
	app.GET("/bankDetail/:username", s.GetBankDetail)
	app.POST("/bankDetail/add", s.AddBankDetail)

	// Withdrawal
	app.POST("/withdrawal/init", s.InitiateWithdrawal)
	app.POST("/withdrawal/complete", s.CompleteWithdrawal)
	app.GET("/withdrawal/:id", s.GetWithdrawalByID)
	app.GET("/withdrawal/all/:status", s.GetWithdrawalsByStatus)

	// Jackpot
	app.POST("/jackpot", s.Jackpot)
	s.router = router
}

func (s *Server) Start(address string) error {
	return s.router.Run(address)
}