package controllers

import (
	"Go-Gin/initializers"
	"Go-Gin/models"

	"github.com/gin-gonic/gin"
)

func PostsCreate(ctx *gin.Context) {
	// get the request body
	var body struct {
		Title string `json:"title"`
		Body  string `json:"body"`
	}
	ctx.Bind(&body)

	// create a new post
	post := models.Post{
		Title: body.Title,
		Body:  body.Body,
	}
	result := initializers.Db.Create(&post)

	if result.Error != nil {
		ctx.JSON(400, gin.H{
			"error": result.Error,
		})
		return
	}
	// save the post
	ctx.JSON(200, gin.H{
		"post": post,
	})
}

func PostsIndex(ctx *gin.Context) {
	var posts []models.Post
	initializers.Db.Find(&posts)

	ctx.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostsShow(ctx *gin.Context) {
	// get the post by id
	var post models.Post
	initializers.Db.First(&post, ctx.Param("id"))

	ctx.JSON(200, gin.H{
		"post": post,
	})
}

func PostUpdate(ctx *gin.Context) {
	// get the post by id
	var post models.Post
	initializers.Db.First(&post, ctx.Param("id"))

	// get the request body
	var body struct {
		Title string `json:"title"`
		Body  string `json:"body"`
	}
	ctx.Bind(&body)

	// update the post
	initializers.Db.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	ctx.JSON(200, gin.H{
		"post": post,
	})
}

func PostDelete(ctx *gin.Context) {
	// get the post by id
	var post models.Post
	initializers.Db.First(&post, ctx.Param("id"))

	// delete the post
	initializers.Db.Delete(&post)

	ctx.JSON(200, gin.H{
		"post": post,
	})
}
