package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Gandalf-Le-Dev/goapi/models"
	"github.com/Gandalf-Le-Dev/goapi/services"
	"github.com/gin-gonic/gin"
)

// swagger embed files

// Posts godoc
// @Summary Get all posts
// @Description Get details of all posts
// @Tags posts
// @Accept  json
// @Produce  json
// @Success 200
// @Router /posts [get]
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

// CreatePost godoc
// @Summary Create a post
// @Description Create a new post
// @Tags posts
// @Accept  json
// @Produce  json
// @Success 200
// @Router /posts [post]
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

// DeletePost godoc
// @Summary Delete a post
// @Description Delete a post by ID
// @Tags posts
// @Accept json
// @Produce json
// @Param id path int true "Post ID"
// @Success 200 {object} map[string]interface{} "message: Post deleted"
// @Failure 500 {object} map[string]interface{} "error: error description"
// @Router /posts/{id} [delete]
func DeletePost(c *gin.Context) {
	id := c.Param("id")
	err := services.DeletePost(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Fatal(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post deleted"})
}

// UpdatePost godoc
// @Summary Update a post
// @Description Update a post by ID
// @Tags posts
// @Accept json
// @Produce json
// @Param id path int true "Post ID"
// @Success 200
// @Failure 400 {object} map[string]interface{} "error: Invalid form submitted"
// @Failure 500 {object} map[string]interface{} "error: error description"
// @Router /posts/{id} [put]
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

// GetPost godoc
// @Summary Get a post by ID
// @Description Get details of a particular post
// @Tags posts
// @Accept  json
// @Produce  json
// @Param id path int true "Post ID"
// @Success 200
// @Router /posts/{id} [get]
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

// EditPost godoc
// @Summary Edit a post
// @Description Get post details for editing by ID
// @Tags posts
// @Accept json
// @Produce json
// @Produce html
// @Param id path int true "Post ID"
// @Success 200
// @Failure 500 {object} map[string]interface{} "error: error description"
// @Router /posts/{id}/edit [get]
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
