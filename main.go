package main

import (
	"github.com/Gandalf-Le-Dev/goapi/controllers"
	"github.com/Gandalf-Le-Dev/goapi/initializers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()
}

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	r.StaticFile("/", "./web/index.html")
	r.LoadHTMLGlob("./web/templates/*.html")

	r.GET("/posts", controllers.Posts)
	r.GET("/posts/:id", controllers.GetPost)
	r.GET("/posts/:id/edit", controllers.EditPost)
	r.POST("/posts", controllers.CreatePost)
	r.DELETE("/posts/:id", controllers.DeletePost)
	r.PUT("/posts/:id", controllers.UpdatePost)

	r.Run()
}
