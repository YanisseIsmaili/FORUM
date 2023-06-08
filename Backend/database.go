package Forum

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

// TABLE USERS //
type Users struct {
	gorm.Model
	Name     string
	Email    string
	Password string
}

// TABLE POSTES //
type Posts struct {
	gorm.Model
	Title    string
	Content  string
	UserID   uint
	User     User
	Comments []Comment
}

// Commentaires //
type Commentaires struct {
	gorm.Model
	Content string
	UserID  uint
	User    User
	PostID  uint
	Post    Post
}

// LocaUser //
type LocationUser struct {
	UserID  uint
	User    user
	Comment []comment
}

// CompteAdmin //
type CompteAdmin struct {
	gorm.Model
	Name     string
	Email    string
	Password string
}

// Fonction DATA

func data() {
	dsn := "localhost:8080"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("erreur de connection à la database")
	}

	// Migration
	db.AutoMigrate(&User{}, &Post{}, &Comment{})

	// Crée USER
	user := User{Name: "Clement Garcia", Email: "clement.garcia@gmail.com", Password: "1234"}
	db.Create(&user)

	// Crée un post
	post := Post{Title: "Mon Premier Post", Content: "Hello, World!", UserID: user.ID}
	db.Create(&post)

	// Crée un commentaire
	comment := Comment{Content: "Cool post sa !", UserID: user.ID, PostID: post.ID}
	db.Create(&comment)

	// Requête avec commentaires
	var result Post
	db.Preload("Commentaire").First(&result, post.ID)

	fmt.Println(result)

	// Je supprime le fichier pour éviter les doublons. //
	os.Remove("nerdz.db")

	log.Println("Creating nerdz.db...")
	file, err := os.Create("nerds.db") // Crée un fichier sqlite
	if err != nil {
		log.Fatal(err.Error())
	}
	file.Close()
	log.Println("nerdz.db created")
}
