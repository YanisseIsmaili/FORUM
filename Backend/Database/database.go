package database

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

// TABLE USERS
type Users struct {
	gorm.Model
	Username string
	Email    string
	Password string
}

// TABLE POSTS
type Posts struct {
	gorm.Model
	TitlePost       string
	ContentCategory string
	UserID          string
	User            string
	Theme           string
	CommentsUser    []Comments `gorm:"foreignkey:PostID"`
	Links           string
	Date            time.Time
	PostsID         string
	// user_picture string
}

// TABLE COMMENTS
type Comments struct {
	gorm.Model
	Content string
	UserID  string
	User    Users
	PostID  string
	Post    Posts
}

func CreateDB(db *gorm.DB) {
	// Création des tables
	db.Migrator().DropTable( &Posts{})
	db.AutoMigrate(&Users{}, &Posts{}, &Comments{})

}

// Fonction pour ajouter un utilisateur
func AddUser(username string, email string, password string, db *gorm.DB) {
	// Ajout de l'utilisateur
	db.Create(&Users{Username: username, Email: email, Password: password})

}

// Fonction pour ajouter un post
func AddPost(title string, content string, theme string, postId string, db *gorm.DB) {
	// Ajout du post

	db.Create(&Posts{TitlePost: title, ContentCategory: content, Theme: theme, PostsID: postId, Date: time.Now()})

}

// Fonction pour ajouter un commentaire
func AddComment(content string, userID string, postID string, db *gorm.DB) {
	// Ajout du commentaire
	db.Create(&Comments{Content: content, UserID: userID, PostID: postID})
}
func GetPostFromBdd(db *gorm.DB) ([]Posts, error) {
	var posts []Posts
	result := db.Find(&posts)

	for _, post := range posts {
		fmt.Println("Title:", post.TitlePost)
		fmt.Println("Content:", post.ContentCategory)
	}

	fmt.Println(result)

	return posts, nil
}

func GetAllPosts(db *gorm.DB) ([]Posts, error) {
	var posts []Posts
	if err := db.Find(&posts).Error; err != nil {
		return nil, err
	}

	return posts, nil
}

// // Fonction pour afficher tous les utilisateurs
// func ShowUsers(db *gorm.DB) {
// 	// Affichage des utilisateurs
// 	var users []Users
// 	db.Find(&users)
// 	fmt.Println(users)
// }

// // Fonction pour afficher tous les posts
// func ShowPosts(db *gorm.DB) {

// 	// Affichage des posts
// 	var posts []Posts
// 	db.Find(&posts)
// 	fmt.Println(posts)
// }

// // Fonction pour afficher tous les commentaires
// func ShowComments(db *gorm.DB) {
// 	// Affichage des commentaires
// 	var comments []Comments
// 	db.Find(&comments)
// 	fmt.Println(comments)
//}
