package service

import (
	"fmt"

	database "Forum/Backend/Database"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetCommentsByPostID(postID string, db *gorm.DB) ([]database.Comments, error) {
	var comments []database.Comments
	result := db.Where("post_id = ?", postID).Find(&comments)
	if result.Error != nil {
		return nil, result.Error
	}
	return comments, nil
}
func ShowPostModal(c *gin.Context, db *gorm.DB) {
	postID := c.Query("postID")
	postIDStr := postID

	comments, err := GetCommentsByPostID(postIDStr, db)
	if err != nil {
		// Gérer l'erreur de récupération des commentaires
		// ...
		return
	}

	for _, comment := range comments {
		fmt.Println("Comment:", comment.Content)
	}
}
