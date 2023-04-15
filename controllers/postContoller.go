package controllers

import (
	"GO_crud/initializers"
	"GO_crud/models"

	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {
	// Получаем данные от requset запроса
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	// Записывает в БД данные от requet
	post := models.Post{Title: body.Title, Body: body.Body}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostIndex(c *gin.Context) {
	// Получаем все посты
	var posts []models.Post
	initializers.DB.Find(&posts)

	c.JSON(200, gin.H{
		"post": posts,
	})
}

func PostShow(c *gin.Context) {
	// Получаем id нужного нам поста
	id := c.Param("id")

	// Получаем один пост по переданному id
	var post models.Post
	initializers.DB.First(&post, id)

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostUpdate(c *gin.Context) {
	// Получаем id нужного нам поста
	id := c.Param("id")

	// Формируем заново структуру поста
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	// Получаем один пост по переданному id
	var post models.Post
	initializers.DB.First(&post, id)

	// Обновляем пост с новыми данными
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostDelete(c *gin.Context) {
	// Получаем id нужного нам поста
	id := c.Param("id")

	// Удаляем пост по полученному id
	initializers.DB.Delete(&models.Post{}, id)

	c.JSON(200, gin.H{
		"message": "Post deleted",
	})
}
