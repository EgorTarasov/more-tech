package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	_ "more-tech/docs"
	"more-tech/internal/api/middleware"
	"more-tech/internal/config"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.mongodb.org/mongo-driver/mongo"
)

type router struct {
	engine      *gin.Engine
	mongoClient *mongo.Client
}

func NewRouter(mongoClient *mongo.Client) router {
	return router{
		engine:      gin.Default(),
		mongoClient: mongoClient,
	}
}

func (r *router) Run(serverPort string) error {
	r.setup()

	return r.engine.Run(":" + serverPort)
}

func (r *router) setup() {
	var domain, mlHost, routeHost string
	if config.Cfg.DockerMode {
		domain = "larek.itatmisis.ru:9999"
		mlHost = "ml"
		routeHost = "larek.itatmisis.ru:10000"
	} else {
		domain = "localhost"
		mlHost = "localhost"
		routeHost = "larek.itatmisis.ru:10000"
	}

	r.engine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}))
	r.engine.Use(middleware.AuthMiddleware(domain))
	r.engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := r.engine.Group("/v1")

	r.setupUserRoutes(v1)
	r.setupDepartmentRoutes(v1, routeHost)
	r.setupTicketRoutes(v1, routeHost)
	r.setupSearchtRoutes(v1, mlHost)
}
