package main

import (
	"fmt"
	"net/http"

	"github.com/Virtual-Dynamic-Labs/DynoV-Open-Metaverse-Cloud/blockchain/docs"
	clients "github.com/Virtual-Dynamic-Labs/DynoV-Open-Metaverse-Cloud/blockchain/internal/clients"
	"github.com/Virtual-Dynamic-Labs/DynoV-Open-Metaverse-Cloud/blockchain/pkg/log"
	"github.com/Virtual-Dynamic-Labs/DynoV-Open-Metaverse-Cloud/blockchain/pkg/storage"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Server struct {
	Database 		*storage.DB
	Router   		*gin.Engine
	Web3Clients 	WebClients
	Logger   		log.Logger
}

type WebClients struct {
	EthClient *clients.EthClient
}

var swagHandler gin.HandlerFunc

func init() {
	swagHandler = ginSwagger.WrapHandler(swaggerFiles.Handler)
}
// @title           DynoV Open Metaverse Cloud Blockchain Service
// @version         1.0
// @description     This is a DynoV Open Metaverse Cloud Blockchain Service.

// @contact.name   Virtual Dynamic Labs Support
// @contact.url    http://www.virtualdynamiclabs.xyz
// @contact.email  michaelzhou@virtualdynamiclabs.xyz
//
// @host      localhost:8080
// @BasePath  /api/alpha
// @schemes   http https
//
// @securityDefinitions.basic  BasicAuth
func main() {
	docs.SwaggerInfo.Title = "DynoV Open Metaverse Cloud - Blockchain Service"

	// create root logger tagged with server version
	logger := log.New("Cloud Blockchain Service")

	db, err := storage.ConnectToDatabase()
	if err != nil {
		panic(fmt.Sprintf("cannot connect to databse because %s", err))
	}

	r := gin.Default()

	clients := WebClients{
		// RPC provider URL should read from configuration file.
		EthClient: clients.CreateEthereumClient("example-url", logger),
	}

	hubServer := Server{
		Database: 	 storage.NewDatabase(db),
		Router:   	 r,
		Web3Clients: clients,
		Logger:   	 logger,
	}

	hubServer.Router = hubServer.setupRouter()

	hubServer.Deploy()
}

func (s *Server) Deploy() {
	if err := s.Router.Run("localhost:8080"); err != nil {
		panic(err)
	}
}

func (s *Server) setupRouter() *gin.Engine {
	r := s.Router

	alpha := r.Group("/api/alpha")
	beta := r.Group("/api/beta")
	v1 := r.Group("/api/v1")

	// Ping test
	alpha.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong with alpha")
	})
	beta.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong with beta")
	})
	v1.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong with v1")
	})

	alpha.GET("/getBlockNumber", func(c *gin.Context) {
		blockNumber := s.Web3Clients.EthClient.GetBlockNumber()
		c.String(http.StatusOK, fmt.Sprintf("Block Number: %d", blockNumber))
	})

	if swagHandler != nil {
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	return r
}