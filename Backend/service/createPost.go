package service

import (
	database "Forum/Backend/database"
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreatePost(c *gin.Context, db *gorm.DB, token string) {

	fmt.Println("____________________CREATE_______________________POST_______________________")

	username, err := GetUsernameFromToken(token)

	if err != nil {
		// Gérer l'erreur de récupération de l'ID d'utilisateur
		fmt.Printf("Erreur de récupération de l'ID d'utilisateur: %s", err)
		// Par exemple, renvoyer une erreur ou effectuer une autre action appropriée
	}

	title := c.PostForm("post-title")
	content := c.PostForm("post-content")
	Theme := c.PostForm("spec")

	post := database.Posts{
		TitlePost:       title,
		ContentCategory: content,
		Theme:           Theme,
	}

	fmt.Println(username)
	fmt.Println(post)

	database.AddPost(post.TitlePost, post.ContentCategory, post.Theme, db)
}

// username := c.PostForm("username")
// 	email := c.PostForm("email")
// 	password := c.PostForm("password")

// 	// Création d'un nouvel utilisateur avec les données récupérées
// 	user := Database.Users{
// 		Username: username,
// 		Email:    email,
// 		Password: password,
// 	}

// 	// Appel à la fonction createDB() pour créer et initialiser l'utilisateur
// 	//Database.CreateDB(db)
// 	Database.AddUser(user.Username, user.Email, user.Password, db)

// 	// Fermeture de la connexion à la base de données

// 	// Réponse au client

// }
