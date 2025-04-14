package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	db "simplebanks/db/sqlc"
	"simplebanks/util"
)

type Server struct {
	store  db.Store
	router *gin.Engine
	config util.Config
}

// thiết lập các tuyến API HTTP
func NewServer(config util.Config, store db.Store) (*Server, error) {

	server := &Server{store: store,
		config: config}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}
	server.SetRouter()
	return server, nil
}

func (server *Server) SetRouter() {
	router := gin.Default()
	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.GetAccount)
	router.GET("/accounts", server.ListAccount)
	router.PUT("/accounts/:id", server.updateAccount)
	router.DELETE("/accounts/:id", server.deleteAccount)
	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error :": err.Error()}
}
