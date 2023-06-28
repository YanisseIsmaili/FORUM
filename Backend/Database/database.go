package database

import (
	"fmt"
	"time"

	"github.com/google/uuid"
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
	PostsID         string `gorm:"column:PostsID"`

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

	db.AutoMigrate(&Users{}, &Posts{}, &Comments{})

}

// Fonction pour ajouter un utilisateur
func AddUser(username string, email string, password string, db *gorm.DB) {
	// Ajout de l'utilisateur
	db.Create(&Users{Username: username, Email: email, Password: password})

}

// Fonction pour ajouter un post
func AddPost(title string, content string, theme string, db *gorm.DB) {
	// Générer un identifiant unique pour le post
	postID := uuid.NewString()

	// Ajout du post
	post := &Posts{
		TitlePost:       title,
		ContentCategory: content,
		Theme:           theme,
		PostsID:         postID,
		Date:            time.Now(),
	}

	db.Create(post)
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
func GetPostByID(postID string, db *gorm.DB) (*Posts, error) {
	var post Posts
	if err := db.Where("posts_id = ?", postID).Preload("CommentsUser").First(&post).Error; err != nil {
		return nil, err
	}

	return &post, nil
}
