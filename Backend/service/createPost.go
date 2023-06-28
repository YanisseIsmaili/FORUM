package service

import (
	database "Forum/Backend/Database"

	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreatePost(c *gin.Context, db *gorm.DB) { //token string, db *gorm.DB) { pour l'username plus tard

	fmt.Println("____________________CREATE_______________________POST_______________________")

	// username, err := GetUsernameFromToken(token)

	// if err != nil {
	// 	// Gérer l'erreur de récupération de l'ID d'utilisateur
	// 	fmt.Printf("Erreur de récupération de l'ID d'utilisateur: %s", err)
	// 	// Par exemple, renvoyer une erreur ou effectuer une autre action appropriée
	// }

	title := c.PostForm("postUserTitle")
	content := c.PostForm("postUserContent")
	Theme := c.PostForm("spec")

	post := database.Posts{
		TitlePost:       title,
		ContentCategory: content,
		Theme:           Theme,
	}

	// Utiliser postId dans votre logique ou pour l'affecter à une autre variable si nécessaire

	//fmt.Println(username)
	fmt.Println(post)

	database.AddPost(post.TitlePost, post.ContentCategory, post.Theme, db)

}
