package service

import (
	"fmt"

	"strconv"

	database "Forum/Backend/Database"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SendComment(c *gin.Context, db *gorm.DB) {
	postIDStr := c.PostForm("postID") // Récupérer l'ID du post à partir des données du formulaire
	postID, err := strconv.ParseUint(postIDStr, 10, 64)
	if err != nil {
		// Gérer l'erreur de conversion de l'ID du post
		// ...
		return
	}
	username := "bot"

	contentComment := c.PostForm("comment")

	comment := database.Comments{
		PostID:  uint(postID), // Convertir l'ID du post en uint
		Content: contentComment,
		UserID:  username, // Assurez-vous que le type de UserID correspond au type de l'ID utilisateur dans la base de données
	}

	// Insérer le commentaire dans la base de données
	result := db.Create(&comment)
	if result.Error != nil {
		// Gérer l'erreur de création du commentaire
		fmt.Printf("Erreur de création du commentaire : %s", result.Error)
		// Par exemple, renvoyer une erreur ou effectuer une autre action appropriée
	}
}

func GetCommentsByPostID(postID uint, db *gorm.DB) ([]database.Comments, error) {
	var comments []database.Comments
	result := db.Where("post_id = ?", postID).Find(&comments)
	if result.Error != nil {
		return nil, result.Error
	}
	return comments, nil
}

func ShowPostModal(c *gin.Context, db *gorm.DB) {
	// Récupérer l'ID du post à partir des paramètres de requête
	postIDStr := c.Query("postID")
	postID, err := strconv.ParseUint(postIDStr, 10, 64)
	if err != nil {
		// Gérer l'erreur de conversion de l'ID du post
		// ...
		return
	}

	// Charger les commentaires du post depuis la base de données
	comments, err := GetCommentsByPostID(uint(postID), db)
	if err != nil {
		// Gérer l'erreur de récupération des commentaires
		// ...
		return
	}

	// Faites quelque chose avec les commentaires chargés, par exemple les afficher dans la console
	for _, comment := range comments {
		fmt.Println("Comment:", comment.Content)
	}
}
