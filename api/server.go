package api

import (
	"fmt"

	db "github.com/fxfrancky/simplebank/db/sqlc"
	"github.com/fxfrancky/simplebank/token"
	"github.com/fxfrancky/simplebank/util"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// Server serves All Http requests for our banking service
type Server struct {
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

// NewServer creates a new HTTP server and setup routing
func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}
	server.setupRouter()
	return server, nil

}

func (server *Server) setupRouter() {
	router := gin.Default()

	// User
	router.POST("/users", server.createUser)
	router.POST("/users/login", server.loginUser)
	// NB Every Request must go througth the authMiddleware

	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))
	// Account
	authRoutes.POST("/accounts", server.createAccount)
	authRoutes.GET("/accounts/:id", server.getAccount)
	authRoutes.GET("/accounts", server.listAccount)

	// Transfer
	authRoutes.POST("/transfers", server.createTransfer)
	// router.GET("/transfers/:id", server.getTransfer)
	// router.GET("/transfers", server.listTransfer)

	// Entry
	router.POST("/entries", server.createEntry)
	router.GET("/entries/:id", server.getEntry)
	router.GET("/entries", server.listEntry)

	server.router = router

}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

// Start runs the HTTP server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
