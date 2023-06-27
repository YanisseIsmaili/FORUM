package main

import (
	database "Forum/Backend/Database"
	service "Forum/Backend/Service"
	"Forum/Backend/controller"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const IndexURL = "/index"

func main() {
	dsn := "NerdMythology.db"
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
		loginService := service.StaticLoginService(ctx, dbConnector)
		var jwtService service.JWTService = service.JWTAuthService()
		var loginController controller.LoginController = controller.LoginHandler(loginService, jwtService)
		token := loginController.Login(ctx)
		if token != "" {
			fmt.Println("Token found")
			ctx.SetCookie("token", token, 3600, "/", "", false, false)
			ctx.Redirect(http.StatusFound, "/index")
			return
		} else {
			fmt.Println("Token not found")
			ctx.Redirect(http.StatusFound, "/login")
			return
		}
	})

	// Middleware pour vérifier le token à chaque requête
	r.Use(func(c *gin.Context) {
		token, err := c.Cookie("token")
		if err != nil || token == "" {
			// Rediriger vers la page de connexion si le token est invalide ou non présent
			c.Redirect(http.StatusFound, "/login")
			c.Abort()
			return
		}
		c.Next()
	})

	r.GET("/logout", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/login")
	})
	r.POST("/login/register", func(c *gin.Context) {
		service.RegisterUser(c, dbConnector)
		c.Redirect(http.StatusFound, "/login")

	})

	r.GET("/createPost", func(c *gin.Context) {
		c.HTML(http.StatusOK, "Post.html", gin.H{})
	})

	r.POST("/sendPost", func(c *gin.Context) {
		token, _ := c.Cookie("token")
		fmt.Println(token)
		service.CreatePost(c, dbConnector, token)
		c.Redirect(http.StatusFound, "/index")

	})

	r.GET("/index", func(c *gin.Context) {
		// Code pour gérer l'accès à la page /index lorsque le token est valide

		// Code pour gérer l'accès à la page /index lorsque le token est valide
		posts, err := database.GetPostFromBdd(dbConnector)
		if err != nil {
			log.Println("Erreur lors de la récupération des posts:", err)
			c.HTML(http.StatusInternalServerError, "error.html", nil)
			return
		}
		c.HTML(http.StatusOK, "index.html", gin.H{
			"Posts": posts, // Passer les posts au modèle HTML
		})

		

	})
	r.Run(":8089")

}
