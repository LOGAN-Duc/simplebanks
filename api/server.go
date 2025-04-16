package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	db "simplebanks/db/sqlc"
	"simplebanks/token"
	"simplebanks/util"
)

type Server struct {
	store      db.Store
	router     *gin.Engine
	tokenMaker token.Maker
	config     util.Config
}

// thiết lập các tuyến API HTTP
func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("token.NewPasetoMaker: %w", err)
	}
	server := &Server{store: store,
		tokenMaker: tokenMaker,
		config:     config}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}
	server.SetRouter()
	return server, nil
}

func (server *Server) SetRouter() {
	router := gin.Default()
	authRouters := router.Group("/").Use(authMiddleware(server.tokenMaker))
	authRouters.POST("/accounts", server.createAccount)
	authRouters.GET("/accounts/:id", server.GetAccount)
	authRouters.GET("/accounts", server.ListAccount)
	authRouters.PUT("/accounts/:id", server.updateAccount)
	authRouters.DELETE("/accounts/:id", server.deleteAccount)
	//user
	router.POST("/users", server.createUser)
	router.POST("/users/login", server.loginUser)
	//transfer
	authRouters.POST("/transfers", server.createTransfer)
	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error :": err.Error()}
}
func (server *Server) GetRouter() *gin.Engine {
	return server.router
}
