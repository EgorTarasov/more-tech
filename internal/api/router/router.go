package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	_ "more-tech/docs"

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
	r.engine.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"*"},
		AllowHeaders: []string{"*"},
	}))
	r.engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := r.engine.Group("/v1")

	r.setupUserRoutes(v1)
	r.setupDepartmentRoutes(v1)
}
