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
func CreateDB() {
	// Connexion à la base de données
<<<<<<< HEAD
	dsn := "Forum:1234@tcp(localhost:8080)/NerdzMethology?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Erreur de connexion à la base de données:", err)
	}
}

// Fonction pour créer les tables
func CreateTables() {
	// Connexion à la base de données
	dsn := "Forum:1234@tcp(localhost:8080)/NerdzMethology?charset=utf8mb4&parseTime=True&loc=Local"
=======
	dsn := "Forum:1234@tcp(:8080)/NerdzMethology?charset=utf8mb4&parseTime=True&loc=Local"
>>>>>>> 70a9d312c3d0995081ed54b777ad7864dc2b592e
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Erreur de connexion à la base de données:", err)
	}

	// Création des tables
	db.AutoMigrate(&Users{}, &Posts{}, &Comments{})
}

// Fonction pour supprimer les tables
func DropTables() {
	// Connexion à la base de données
	dsn := "Forum:1234@tcp(localhost:8080)/NerdzMethology?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Erreur de connexion à la base de données:", err)
	}

	// Suppression des tables
	db.Migrator().DropTable(&Users{}, &Posts{}, &Comments{})
}

// Fonction pour ajouter un utilisateur
func AddUser(username string, email string, password string) {
	// Connexion à la base de données
	dsn := "Forum:1234@tcp(localhost:8080)/NerdzMethology?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Erreur de connexion à la base de données:", err)
	}

	// Ajout de l'utilisateur
	db.Create(&Users{Username: username, Email: email, Password: password})
}

// Fonction pour ajouter un post
func AddPost(title string, content string, userID uint) {
	// Connexion à la base de données
	dsn := "Forum:1234@tcp(localhost:8080)/NerdzMethology?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Erreur de connexion à la base de données:", err)
	}

	// Ajout du post
	db.Create(&Posts{Title: title, Content: content, UserID: userID})
}

// Fonction pour ajouter un commentaire
func AddComment(content string, userID uint, postID uint) {
	// Connexion à la base de données
	dsn := "Forum:1234@tcp(localhost:8080)/NerdzMethology?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Erreur de connexion à la base de données:", err)
	}

	// Ajout du commentaire
	db.Create(&Comments{Content: content, UserID: userID, PostID: postID})
}

// Fonction pour afficher tous les utilisateurs
func ShowUsers() {
	// Connexion à la base de données
	dsn := "Forum:1234@tcp(localhost:8080)/NerdzMethology?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Erreur de connexion à la base de données:", err)
	}

	// Affichage des utilisateurs
	var users []Users
	db.Find(&users)
	fmt.Println(users)
}

// Fonction pour afficher tous les posts
func ShowPosts() {
	// Connexion à la base de données
	dsn := "Forum:1234@tcp(localhost:8080)/NerdzMethology?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Erreur de connexion à la base de données:", err)
	}

	// Affichage des posts
	var posts []Posts
	db.Find(&posts)
	fmt.Println(posts)
}

// Fonction pour afficher tous les commentaires
func ShowComments() {
	// Connexion à la base de données
	dsn := "Forum:1234@tcp(localhost:8080)/NerdzMethology?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Erreur de connexion à la base de données:", err)
	}

	// Affichage des commentaires
	var comments []Comments
	db.Find(&comments)
	fmt.Println(comments)
}

// Fonction pour afficher un utilisateur
func ShowUser(id uint) {
	// Connexion à la base de données
	dsn := "Forum:1234@tcp(localhost:8080)/NerdzMethology?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// Affichage d'un utilisateur
	var user Users
	db.First(&user, id)
	fmt.Println(user)
}

// Fonction pour afficher un post
func ShowPost(id uint) {
	// Connexion à la base de données
	dsn := "Forum:1234@tcp(localhost:8080)/NerdzMethology?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// Affichage d'un post
	var post Posts
	db.First(&post, id)
	fmt.Println(post)
}

// Fonction pour afficher un commentaire
func ShowComment(id uint) {
	// Connexion à la base de données
	dsn := "Forum:1234@tcp(localhost:8080)/NerdzMethology?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// Affichage d'un commentaire
	var comment Comments
	db.First(&comment, id)
	fmt.Println(comment)
}

// Fonction pour modifier un utilisateur
func UpdateUser(id uint, username string, email string, password string) {
	// Connexion à la base de données
	dsn := "Forum:1234@tcp(localhost:8080)/NerdzMethology?parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// Modification de l'utilisateur
	var user Users
	db.First(&user, id)
	user.Username = username
	user.Email = email
	user.Password = password
	db.Save(&user)
}

// Fonction pour modifier un post
func UpdatePost(id uint, title string, content string, userID uint) {
	// Connexion à la base de données
	dsn := "Forum:1234@tcp(localhost:8080)/NerdzMethology?parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// Modification du post
	var post Posts
	db.First(&post, id)
	post.Title = title
	post.Content = content
	post.UserID = userID
	db.Save(&post)
}

// Fonction pour modifier un commentaire	// Connexion à la base de données
func UpdateComment(id uint, content string, userID uint, postID uint) {
	dsn := "Forum:1234@tcp(localhost:8080)/NerdzMethology?parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// Modification du commentaire
	var comment Comments
	db.First(&comment, id)
	comment.Content = content
	comment.UserID = userID
	comment.PostID = postID
	db.Save(&comment)
}

// Fonction pour supprimer un utilisateur
func DeleteUser(id uint) {
	// Connexion à la base de données
	dsn := "Forum:1234@tcp(localhost:8080)/NerdzMethology?parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// Suppression de l'utilisateur
	var user Users
	db.First(&user, id)
	db.Delete(&user)
}

// Fonction pour supprimer un post
func DeletePost(id uint) {
	// Connexion à la base de données
	dsn := "Forum:1234@tcp(localhost:8080)/NerdzMethology?parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// Suppression du post
	var post Posts
	db.First(&post, id)
	db.Delete(&post)
}

// Fonction pour supprimer un commentaire
func DeleteComment(id uint) {
	// Connexion à la base de données
	dsn := "Forum:1234@tcp(localhost:8080)/NerdzMethology?parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// Suppression du commentaire
	var comment Comments
	db.First(&comment, id)
	db.Delete(&comment)
}

// Fonction pour supprimer tous les utilisateurs
func DeleteAllUsers() {
	// Connexion à la base de données
	dsn := "Forum:1234@tcp(localhost:8080)/NerdzMethology?parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// Suppression de tous les utilisateurs
	var users []Users
	db.Find(&users)
	db.Delete(&users)
}

// Fonction pour supprimer tous les posts
func DeleteAllPosts() {

	// Connexion à la base de données
	dsn := "Forum:1234@tcp(localhost:8080)/NerdzMethology?parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// Suppression de tous les posts
	var posts []Posts
	db.Find(&posts)
	db.Delete(&posts)
}

// Fonction pour supprimer tous les commentaires
func DeleteAllComments() {
	
	// Connexion à la base de données
	dsn := "Forum:1234@tcp(localhost:8080)/NerdzMethology?parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// Suppression de tous les commentaires
	var comments []Comments
	db.Find(&comments)
	db.Delete(&comments)
}

// Fonction pour supprimer la base de données
func DeleteDB() {
	
	// Suppression de la base de données
	os.Remove("NerdzMethology")
}

// Fonction pour supprimer toutes les tables
func DeleteAllTables() {
	
	// Suppression de toutes les tables
	os.Remove("Users")
	os.Remove("Posts")
	os.Remove("Comments")
}

// Fonction pour supprimer toutes les données
func DeleteAllData() {
	
	// Suppression de toutes les données
	os.Remove("Users")
	os.Remove("Posts")
	os.Remove("Comments")
}
