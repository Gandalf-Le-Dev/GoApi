package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Gandalf-Le-Dev/goapi/models"
	"github.com/Gandalf-Le-Dev/goapi/services"
	"github.com/gin-gonic/gin"
)

func Posts(c *gin.Context) {
	var posts models.PostSlice
	var err error
	posts, err = services.Posts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Fatal(err)
		return
	}

	// Check if the client accepts HTML responses
	if c.GetHeader("Accept") == "text/html" {
		fmt.Print(posts)
		c.HTML(http.StatusOK, "posts.html", gin.H{"posts": posts})
		return
	}

	c.JSON(http.StatusOK, posts)
}

func CreatePost(c *gin.Context) {
	post, err := services.CreatePost(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Fatal(err)
		return
	}

	if c.GetHeader("Accept") == "text/html" {
		c.HTML(http.StatusOK, "post.html", gin.H{"post": post})
		return
	}

	c.JSON(http.StatusOK, post)
}

func DeletePost(c *gin.Context) {
	id := c.Param("id")
	err := services.DeletePost(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Fatal(err)
		return
	}

	// if c.GetHeader("Accept") == "text/html" {
	// 	c.Header("Content-Type", "text/html")
	// 	c.String(http.StatusOK, "")
	// 	return
	// }

	c.JSON(http.StatusOK, gin.H{"message": "Post deleted"})
}

func UpdatePost(c *gin.Context) {
	id := c.Param("id")
	updatedPost, err := services.UpdatePost(id, c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Fatal(err)
		return
	}

	if err := c.ShouldBind(&updatedPost); err != nil {
		c.String(http.StatusBadRequest, "Invalid form submitted")
		return
	}

	if c.GetHeader("Accept") == "text/html" {
		c.HTML(http.StatusOK, "post.html", gin.H{"post": updatedPost})
		return
	}

	c.JSON(http.StatusOK, gin.H{"post": updatedPost})
}

func GetPost(c *gin.Context) {
	id := c.Param("id")
	post, err := services.GetPost(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Fatal(err)
		return
	}

	if c.GetHeader("Accept") == "text/html" {
		c.HTML(http.StatusOK, "post.html", gin.H{"post": post})
		return
	}

	c.JSON(http.StatusOK, post)
}

func EditPost(c *gin.Context) {
	id := c.Param("id")
	post, err := services.GetPost(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Fatal(err)
		return
	}

	if c.GetHeader("Accept") == "text/html" {
		c.HTML(http.StatusOK, "edit_post.html", gin.H{"post": post})
		return
	}

	c.JSON(http.StatusOK, post)
}
