package service

import (
	database "Forum/Backend/Database"
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateComment(c *gin.Context, db *gorm.DB) { //token string, db *gorm.DB) { pour l'username plus tard

	contentComment := c.PostForm("comment")

	comments := database.Comments{
		TitlePost:       title,
		ContentCategory: content,
		Theme:           Theme,
	}

	//fmt.Println(username)



}
