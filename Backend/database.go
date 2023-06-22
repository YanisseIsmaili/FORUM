package Forum

import (
	"fmt"
	"log"
	"os"
	"github.com/joho/godotenv"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string
	Email    string
	Password string
}

type Post struct {
	gorm.Model
	Title    string
	Content  string
	UserID   uint
	User     User
	Comments []Comment
}

type Comment struct {
	gorm.Model
	Content string
	UserID  uint
	User    User
	PostID  uint
	Post    Post
}

func establishDBConnection() (*gorm.DB, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erreur lors du chargement du fichier .env:", err)
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("Erreur de connexion à la base de données: %v", err)
	}

	return db, nil
}

func createDBTables() error {
	db, err := establishDBConnection()
	if err != nil {
		return err
	}

	err = db.AutoMigrate(&User{}, &Post{}, &Comment{})
	if err != nil {
		return fmt.Errorf("Erreur lors de la création des tables: %v", err)
	}

	return nil
}

func addUser(username, email, password string) error {
	db, err := establishDBConnection()
	if err != nil {
		return err
	}

	user := User{
		Username: username,
		Email:    email,
		Password: password,
	}

	err = db.Create(&user).Error
	if err != nil {
		return fmt.Errorf("Erreur lors de la création de l'utilisateur: %v", err)
	}

	return nil
}
