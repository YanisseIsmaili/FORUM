package main

import (
	database "Forum/Backend/Database"
	service "Forum/Backend/Service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const IndexURL = "/index"

func main() {
	service.Test()
	dsn := "forum.db"
	dbConnector, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Erreur de connexion à la base de données:", err)
	}

	// Création et initialisation de la base de données
	database.CreateDB(dbConnector)

	// création route par default
	r := gin.Default()
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Printf("endpoint %v %v %v %v\n", httpMethod, absolutePath, handlerName, nuHandlers)
	}
	r.StaticFile("/styles.css", "./Frontend/styles.css") // permet de link le fichier css
	r.LoadHTMLGlob("Frontend/*")                         // permet d'allez chercher

	// crée une route /get pour afficher la page
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "log&Signup.html", gin.H{})
	})

	// permet de vérifier les informatinos du méthod post
	r.POST("/login", func(ctx *gin.Context) {

	})

	r.GET("/logout", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/login")
	})
	r.POST("/login/register", func(c *gin.Context) {

		service.RegisterUser(c, dbConnector)
		c.Redirect(http.StatusFound, "/login")

	})

	r.Run(":8089")

}
