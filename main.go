package main

import (
	Forum "Forum/Backend"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	dsn := "gorm.db"
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Erreur de connexion à la base de données:", err)
	}

	// route par default :
	r := gin.Default()
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Printf("endpoint %v %v %v %v\n", httpMethod, absolutePath, handlerName, nuHandlers)
	}
	r.StaticFile("/styles.css", "./Frontend/styles.css") // permet de link le fichier css
	r.LoadHTMLGlob("Frontend/*")                         // permet d'allez chercher
	r.GET("/loging", func(c *gin.Context) {
		c.HTML(http.StatusOK, "log&Signup.html", gin.H{})
	})
	r.POST("/loging", func(c *gin.Context) {
		
		username := c.PostForm("username")
		email := c.PostForm("email")
		password := c.PostForm("password")

		// Création d'un nouvel utilisateur avec les données récupérées
		user := Forum.Users{
			Username: username,
			Email:    email,
			Password: password,
		}

		// Appel à la fonction createDB() pour créer et initialiser la base de données
		Forum.CreateDB(db)

		// Connexion à la base de données

		// Insertion de l'utilisateur dans la base de données
		err = db.Create(&user).Error
		if err != nil {
			log.Fatal("Erreur lors de l'insertion de l'utilisateur dans la base de données:", err)
		}

		// Fermeture de la connexion à la base de données
		sqlDB, err := db.DB()
		if err != nil {
			log.Fatal("Erreur lors de la fermeture de la connexion à la base de données:", err)
		}
		sqlDB.Close()

		// Réponse au client
		c.String(http.StatusOK, "Données insérées dans la base de données.")
	})

	r.Run()

	Forum.CreateDB(db)

	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

}
