package Forum

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
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
func createDB() {
	// Connexion à la base de données
	dsn := "Forum:1234@tcp(127.0.0.1:8080)/NerdzMethology?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Erreur de connexion à la base de données:", err)
	}

	// Création des tables s'ils n'existent pas déjà
	err = db.AutoMigrate(&Users{}, &Posts{}, &Comments{})
	if err != nil {
		log.Fatal("Erreur lors de la création des tables:", err)
	}

	// Création d'un utilisateur
	user := Users{Username: "Clement Garcia", Email: "clement.garcia@gmail.com", Password: "1234"}
	db.Create(&user)

	// Création d'un post
	post := Posts{Title: "Mon Premier Post", Content: "Hello, World!", UserID: user.ID}
	db.Create(&post)

	// Création d'un commentaire
	comment := Comments{Content: "Cool post !", UserID: user.ID, PostID: post.ID}
	db.Create(&comment)

	// Récupération d'un post avec ses commentaires
	var result Posts
	db.Preload("Comments").First(&result, post.ID)

	fmt.Println(result)

	// Fermeture de la connexion à la base de données
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Erreur lors de la fermeture de la connexion à la base de données:", err)
	}
	sqlDB.Close()
	// Suppression du fichier de base de données pour éviter les doublons
	err = os.Remove("nerdz.db")
	if err != nil {
		log.Fatal("Erreur lors de la suppression du fichier de base de données:", err)
	}
}
