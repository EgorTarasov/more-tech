package main

import (
	"context"
	"log"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Department struct {
	Id           int    `json:"id"`
	BiskvitId    string `json:"Biskvit_id"`
	ShortName    string `json:"shortName"`
	Address      string `json:"address"`
	City         string `json:"city"`
	ScheduleFl   string `json:"scheduleFl"`
	ScheduleJurL string `json:"scheduleJurL"`
	Special      struct {
		VipZone   int `json:"vipZone"`
		VipOffice int `json:"vipOffice"`
		Ramp      int `json:"ramp"`
		Person    int `json:"person"`
		Juridical int `json:"juridical"`
		Prime     int `json:"Prime"`
	} `json:"special"`
	Coordinates struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	} `json:"coordinates"`
}

type Server struct {
	// db *sql.DB
	mongo *mongo.Client
}

func (s *Server) GetMongo() *mongo.Client {
	return s.mongo
}

func (s *Server) GetDepartmentsCollection() *mongo.Collection {
	return s.mongo.Database("dev").Collection("departments")
}

func (s *Server) getDepartmentById(c *gin.Context) {
	// db := s.GetDb()
	id, _ := strconv.Atoi(c.Param("id"))
	log.Println(id)
	department := s.GetDepartmentsCollection()
	var result bson.M
	department.FindOne(context.TODO(), bson.D{{"id", id}}).Decode(&result)
	log.Println(result)

	c.JSON(200, result)
}

type DepartmentRangeRequest struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Radius    float64 `json:"radius"`
}

type Location struct {
	Latitude  float64
	Longitude float64
}

func (s *Server) getDepartmentByRange(c *gin.Context) {
	var request DepartmentRangeRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	departments := s.GetDepartmentsCollection()
	var results []*Department

	query := bson.M{
		"location": bson.M{
			"$geoWithin": bson.M{
				"$centerSphere": bson.A{
					[]float64{request.Longitude, request.Latitude},
					request.Radius / 6380.752, // convert radius from miles to radians
				},
			},
		},
	}

	cur, err := departments.Find(context.TODO(), query)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}
	cur.All(context.Background(), &results)
	c.JSON(200, results)

}

func mainn() {
	r := gin.Default()

	// loading .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	uri := os.Getenv("MONGO_URI")

	// connecting to mongodb
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	// db, err := sql.Open("sqlite3", "dsn")
	// if err != nil {
	// 	log.Fatalf("cant open db: %+v", err)
	// }

	s := Server{
		mongo: client,
	}

	r.GET("/departments/:id", s.getDepartmentById)
	r.POST("/departments", s.getDepartmentByRange)
	r.Run()
}
