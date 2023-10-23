package main

import (
	"github.com/Gandalf-Le-Dev/goapi/controllers"
	docs "github.com/Gandalf-Le-Dev/goapi/docs"
	"github.com/Gandalf-Le-Dev/goapi/initializers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Go API
// @version         1.0
// @description     A blog API to learn Go.

// @host      localhost:3000
// @BasePath  /api/v1

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()
}

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	r.StaticFile("/", "./web/index.html")
	r.LoadHTMLGlob("./web/templates/*.html")

	// Swagger
	docs.SwaggerInfo.BasePath = "/api"
	api := r.Group("/api")
	{
		api.GET("/posts", controllers.Posts)
		api.GET("/posts/:id", controllers.GetPost)
		api.GET("/posts/:id/edit", controllers.EditPost)
		api.POST("/posts", controllers.CreatePost)
		api.DELETE("/posts/:id", controllers.DeletePost)
		api.PUT("/posts/:id", controllers.UpdatePost)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run()
}
