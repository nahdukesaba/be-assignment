package server

import (
	"github.com/gin-gonic/gin"
	"github.com/nahdukesaba/be-assignment/handlers"
	"github.com/nahdukesaba/be-assignment/repo"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Server struct {
	router *gin.Engine
	db     *repo.DB
}

func NewServer(router *gin.Engine, db *repo.DB) *Server {
	return &Server{
		router: router,
		db:     db,
	}
}

//	@title			Swagger Example API
//	@version		1.0
//	@description	This is a sample server celler server.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:9000
//	@BasePath	/api

//	@securityDefinitions.basic	BasicAuth

//	@externalDocs.description	OpenAPI
//	@externalDocs.url			https://swagger.io/resources/open-api/
func (s *Server) RegisterRoutes() {
	userHandler := handlers.NewUserHandler(s.db)
	paymentHandler := handlers.NewPaymentHandler(s.db)

	userGroup := s.router.Group("/api/user")
	{
		userGroup.POST("/login", userHandler.Login)
		userGroup.POST("/register", userHandler.Register)
	}

	paymentGroup := s.router.Group("/api/payment")
	{
		paymentGroup.POST("/send", paymentHandler.Send)
		paymentGroup.POST("/withdraw", paymentHandler.Withdraw)
	}

	accountGroup := s.router.Group("/api/accounts")
	{
		accountGroup.GET("/", paymentHandler.GetAccountsUser)
		accountGroup.GET("/:account_id", paymentHandler.GetTransactionsAccount)
	}

	s.router.GET("/api/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func (s *Server) Run() {
	s.router.Run(":9000")
}
