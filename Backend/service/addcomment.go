package service

import (
	"fmt"
	"strconv"

	database "Forum/Backend/Database"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetCommentsByPostID(postID uint, db *gorm.DB) ([]database.Comments, error) {
	var comments []database.Comments
	result := db.Where("post_id = ?", postID).Find(&comments)
	if result.Error != nil {
		return nil, result.Error
	}
	return comments, nil
}

func ShowPostModal(c *gin.Context, db *gorm.DB) {
	postIDStr := c.Query("postID")
	postID, err := strconv.ParseUint(postIDStr, 10, 64)
	if err != nil {
		// Gérer l'erreur de conversion de l'ID du post
		// ...
		return
	}

	comments, err := GetCommentsByPostID(uint(postID), db)
	if err != nil {
		// Gérer l'erreur de récupération des commentaires
		// ...
		return
	}

	for _, comment := range comments {
		fmt.Println("Comment:", comment.Content)
	}
}
func SendComment(postID string, comment string, dbConnector *gorm.DB) error {
	// Créez une instance du modèle Comment avec les données du commentaire
	newComment := database.Comments{
		PostID:  postID,
		Content: comment,
	}

	// Insérez le commentaire dans la base de données en utilisant GORM
	if err := dbConnector.Create(&newComment).Error; err != nil {
		return err
	}

	return nil
}
