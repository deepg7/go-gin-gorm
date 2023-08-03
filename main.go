package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/deepg7/go-gin-gorm/controllers"
	_ "github.com/deepg7/go-gin-gorm/docs"
	"github.com/deepg7/go-gin-gorm/models"

	"github.com/gin-gonic/gin"
	"github.com/go-co-op/gocron"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title 			Gin Swagger Example API
// @version 		1.0
// @description 	This is a sample server server.

// @host 			localhost:3000
// @BasePath 		/
// @schemes 		http
func main() {

	godotenv.Load()
	opt, err := redis.ParseURL(os.Getenv("redis_url"))
	if err != nil {
		panic(err)
	}

	client := redis.NewClient(opt)
	ctx := context.Background()

	// er := client.Set(ctx, "foo", "bar", 0).Err()
	// if er != nil {
	// 	panic(er)
	// }

	val, errr := client.Get(ctx, "foo").Result()
	if errr != nil {
		panic(errr)
	}
	fmt.Println("foo", val)
	// cron support
	s := gocron.NewScheduler(time.UTC)
	s.Every(100000).Seconds().Do(func() {
		fmt.Println("John Doe")
	})
	s.StartAsync()

	r := gin.Default()

	models.ConnectDatabase()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})

	r.GET("/books", controllers.FindBooks)

	r.POST("/book", controllers.CreateBook)

	r.GET("/book/:id", controllers.FindBook)

	r.PATCH("/book/:id", controllers.UpdateBook)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run()
}
