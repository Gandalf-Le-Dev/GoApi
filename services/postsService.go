package services

import (
	"github.com/Gandalf-Le-Dev/goapi/models"
	"github.com/Gandalf-Le-Dev/goapi/repositories"
	"github.com/Gandalf-Le-Dev/goapi/utils"
	"github.com/gin-gonic/gin"
)

type PostBody struct {
	Title   string
	Content string
}

func Posts() (models.PostSlice, error) {
	return repositories.Posts()
}

func CreatePost(c *gin.Context) (models.Post, error) {
	var post models.Post
	var postBody PostBody

	bindErr := c.Bind(&postBody)
	utils.PanicIfErr(bindErr)

	post.Title = postBody.Title
	post.Content = postBody.Content
	post.UserID = 1

	post, err := repositories.CreatePost(post)
	if err != nil {
		return post, err
	}

	return post, nil
}

func DeletePost(id string) error {
	idAsInt := utils.StringToInt(id)
	return repositories.DeletePost(idAsInt)
}

func UpdatePost(id string, c *gin.Context) (models.Post, error) {
	var post models.Post
	var postBody PostBody
	idAsInt := utils.StringToInt(id)

	bindErr := c.Bind(&postBody)
	utils.PanicIfErr(bindErr)

	post.Title = postBody.Title
	post.Content = postBody.Content

	return repositories.UpdatePost(idAsInt, post)
}

func GetPost(id string) (models.Post, error) {
	idAsInt := utils.StringToInt(id)
	return repositories.GetPostById(idAsInt)
}
