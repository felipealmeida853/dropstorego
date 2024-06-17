package main

import (
	"context"
	"dropstore/config"
	"dropstore/controllers"
	"dropstore/routes"
	"dropstore/services"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	server      *gin.Engine
	ctx         context.Context
	mongoClient *mongo.Client
	redisClient *redis.Client

	userService         services.UserService
	UserController      controllers.UserController
	UserRouteController routes.UserRouteController

	authCollection      *mongo.Collection
	authService         services.AuthService
	AuthController      controllers.AuthController
	AuthRouteController routes.AuthRouteController
)

func init() {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("Error loading config", err)
	}

	ctx = context.Background()

	mongoConn := options.Client().ApplyURI(config.MongoDBURI)
	mongoClient, err := mongo.Connect(ctx, mongoConn)
	if err != nil {
		log.Fatal("Error connecting MongoDB", err)
	}

	if err := mongoClient.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal("Error pingin MongoDB", err)
	}

	fmt.Println("Connected to MongoDB with success")

	redisClient = redis.NewClient(&redis.Options{
		Addr:     config.RedisURI,
		Password: "",
		DB:       0,
	})

	if _, err = redisClient.Ping(ctx).Result(); err != nil {
		log.Fatal("Error pinging Redis", err)
	}

	err = redisClient.Set(ctx, "healthcheck", "Golang API with Redis and Mongo, OK!", 0).Err()
	if err != nil {
		log.Fatal("Error setting key in redis")
	}

	fmt.Println("Connected to Redis with success")

	authCollection = mongoClient.Database("go").Collection("users")
	userService = services.NewUserServiceImpl(authCollection, ctx)
	authService = services.NewAuthService(authCollection, ctx)
	AuthController = controllers.NewAuthController(authService, userService, ctx, authCollection)
	AuthRouteController = routes.NewAuthRouteController(AuthController)

	UserController = controllers.NewUserController(userService)
	UserRouteController = routes.NewRouteUserController(UserController)

	server = gin.Default()

}

func main() {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("Error loading config", err)
	}
	defer mongoClient.Disconnect(ctx)
	go startGRPCServer(config)
	startGinServer(config)
}

func startGinServer(config config.Config) {
	value, err := redisClient.Get(ctx, "healthcheck").Result()
	if err == redis.Nil {
		fmt.Println("Key healthcheck doesnt exist")
	} else if err != nil {
		panic(err)
	}

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{config.Origin}
	corsConfig.AllowCredentials = true

	server.Use(cors.New(corsConfig))

	router := server.Group("api")
	router.GET("healthcheck", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": value})
	})

	AuthRouteController.AuthRoute(router, userService)
	UserRouteController.UserRoute(router, userService)

	log.Fatal(server.Run(":" + config.Port))
}

func startGRPCServer(config config.Config) {
	fmt.Printf("GRPC %v \n", config)

}
