package router

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type router struct {
	engine *gin.Engine
	mongoClient *mongo.Client
}

func NewRouter(mongoClient *mongo.Client) router {
	return router{
		engine: gin.Default(),
		mongoClient: mongoClient,
	}
}

func (r *router) Run(serverPort string) error {
	r.setup()

	return r.engine.Run(":" + serverPort)
	
}

func (r *router) setup() {
	v1 := r.engine.Group("/v1")

	r.setupUserRoutes(v1)
	r.setupDepartmentRoutes(v1)
}
