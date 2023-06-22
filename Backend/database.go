package Forum

import (
	"fmt"


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
	Title    string
	Content  string
	UserID   uint
	User     Users
	Comments []Comments
}

// TABLE COMMENTS
type Comments struct {
	gorm.Model
	Content string
	UserID  uint
	User    Users
	PostID  uint
	Post    Posts
}

// Fonction pour créer et initialiser la base de données
func CreateDB(db *gorm.DB) {
	// Création des tables
	db.AutoMigrate(&Users{}, &Posts{}, &Comments{})
}

// Fonction pour ajouter un utilisateur
func AddUser(username string, email string, password string, db *gorm.DB) {
	// Ajout de l'utilisateur
	db.Create(&Users{Username: username, Email: email, Password: password})
}

// Fonction pour ajouter un post
func AddPost(title string, content string, userID uint, db *gorm.DB) {
	// Ajout du post
	db.Create(&Posts{Title: title, Content: content, UserID: userID})
}

// Fonction pour ajouter un commentaire
func AddComment(content string, userID uint, postID uint, db *gorm.DB) {
	// Ajout du commentaire
	db.Create(&Comments{Content: content, UserID: userID, PostID: postID})
}

// Fonction pour afficher tous les utilisateurs
func ShowUsers(db *gorm.DB) {
	// Affichage des utilisateurs
	var users []Users
	db.Find(&users)
	fmt.Println(users)
}

// Fonction pour afficher tous les posts
func ShowPosts(db *gorm.DB) {

	// Affichage des posts
	var posts []Posts
	db.Find(&posts)
	fmt.Println(posts)
}

// Fonction pour afficher tous les commentaires
func ShowComments(db *gorm.DB) {
	// Affichage des commentaires
	var comments []Comments
	db.Find(&comments)
	fmt.Println(comments)
}
