package api

import (
	"fmt"
	db "simplebank/db/sqlc"

    "github.com/gin-gonic/gin/binding"
    "github.com/go-playground/validator/v10"
	"github.com/gin-gonic/gin"
)

// Server struct -> serves all HTTP requests for our banking service
// router field is a router of type gin.Engine -> will help us send each API request to the correct handler for processing
type Server struct {
	store db.Store
	router *gin.Engine
}

// NewServer takes a db.Store as input, and return a Server.
// This function create a new Server instance, and setup all HTTP API routes for our service on that server
func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	// TODO: add routes to router
	router.POST("accounts", server.createAccount)

	// get accounts by ID
	router.GET("/accounts/:id", server.getAccount)

	// get list accounts
	router.GET("/accounts", server.listAccounts)

	// transfer money
	router.POST("/transfers", server.createTransfer)

	// create User
	router.POST("/users", server.createUser)

	server.router = router
	return server
}


// Start HTTP server
func (server *Server) Start(address string) error {
	fmt.Println("Server is running at address:", address)
	return server.router.Run(address)
}


