package controllers

import (
	"log"
	"net/http"

	"github.com/Gandalf-Le-Dev/goapi/models"
	"github.com/Gandalf-Le-Dev/goapi/services"
	"github.com/Gandalf-Le-Dev/goapi/utils/htmlFormatter"
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
		html := htmlFormatter.Posts(posts)
		c.Header("Content-Type", "text/html")
		c.String(http.StatusOK, html)
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
		html := htmlFormatter.Post(post)
		c.Header("Content-Type", "text/html")
		c.String(http.StatusOK, html)
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

	if c.GetHeader("Accept") == "text/html" {
		c.Header("Content-Type", "text/html")
		c.String(http.StatusOK, "")
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post deleted"})
}

func UpdatePost(c *gin.Context) {
	id := c.Param("id")
	post, err := services.UpdatePost(id, c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Fatal(err)
		return
	}

	if c.GetHeader("Accept") == "text/html" {
		html := htmlFormatter.Post(post)
		c.Header("Content-Type", "text/html")
		c.String(http.StatusOK, html)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post updated"})
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
		html := htmlFormatter.EditPost(post)
		c.Header("Content-Type", "text/html")
		c.String(http.StatusOK, html)
		return
	}

	c.JSON(http.StatusOK, post)
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
		html := htmlFormatter.Post(post)
		c.Header("Content-Type", "text/html")
		c.String(http.StatusOK, html)
		return
	}

	c.JSON(http.StatusOK, post)
}
