package main

import (
	"backend_test_sharing_vision/api/controller"
	"backend_test_sharing_vision/api/repository"
	"backend_test_sharing_vision/api/routes"
	"backend_test_sharing_vision/api/service"
	"backend_test_sharing_vision/infrastructure"
	"backend_test_sharing_vision/models"

	"github.com/gin-contrib/cors"
)

func init() {
	infrastructure.LoadEnv()
}

func main() {
	router := infrastructure.NewGinRouter()                     //router has been initialized and configured
	db := infrastructure.NewDatabase()                          // databse has been initialized and configured
	postRepository := repository.NewPostRepository(db)          // repository are being setup
	postService := service.NewPostService(postRepository)       // service are being setup
	postController := controller.NewPostController(postService) // controller are being set up
	postRoute := routes.NewPostRoute(postController, router)    // post routes are initialized
	postRoute.Setup()                                           // post routes are being setup

	db.DB.AutoMigrate(&models.Post{}) // migrating Post model to datbase table
	router.Gin.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},  // Ganti dengan domain Anda
		AllowMethods:     []string{"POST"},                   // Izinkan metode POST
		AllowHeaders:     []string{"Origin", "Content-Type"}, // Izinkan header yang dibutuhkan
		AllowCredentials: true,
	}))
	router.Gin.Run(":8080") //server started on 8000 port
}
