package repositories

import (
	"context"

	"github.com/Gandalf-Le-Dev/goapi/initializers"
	"github.com/Gandalf-Le-Dev/goapi/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func Posts() (models.PostSlice, error) {
	var posts models.PostSlice
	var err error
	posts, err = models.Posts().All(context.Background(), initializers.DB)
	if err != nil {
		return posts, err
	}

	return posts, nil
}

func CreatePost(post models.Post) (models.Post, error) {
	err := post.Insert(context.Background(), initializers.DB, boil.Infer())
	if err != nil {
		return post, err
	}

	return post, nil
}

func DeletePost(id int) error {
	_, err := models.Posts(models.PostWhere.PostID.EQ(id)).DeleteAll(context.Background(), initializers.DB)
	if err != nil {
		return err
	}

	return nil
}

func UpdatePost(id int, postBody models.Post) (models.Post, error) {
	post, findErr := models.FindPost(context.Background(), initializers.DB, id)
	if findErr != nil {
		return *post, findErr
	}

	post.Title = postBody.Title
	post.Content = postBody.Content

	_, updateErr := post.Update(context.Background(), initializers.DB, boil.Infer())
	if updateErr != nil {
		return *post, updateErr
	}

	return *post, nil
}

func GetPostById(id int) (models.Post, error) {
	var post *models.Post
	var err error
	post, err = models.FindPost(context.Background(), initializers.DB, id)
	if err != nil {
		return *post, err
	}

	return *post, nil
}
