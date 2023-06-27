package main

import (
	"fmt"
	"log"
	"net/http"

	database "Forum/Backend/Database"
	service "Forum/Backend/Service"
	"Forum/Backend/controller"

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

	r.StaticFile("/styles.css", "./Frontend/styles.css") // permet de lier le fichier CSS
	r.StaticFile("/style.css", "./Frontend/style.css")
	r.LoadHTMLGlob("Frontend/*") // permet d'aller chercher les modèles HTML
	r.POST("/register", func(c *gin.Context) {
		service.RegisterUser(c, dbConnector)
		c.HTML(http.StatusOK, "log&Signup.html", gin.H{})
	})
	r.GET("/authentificationfield", func(c *gin.Context) {
		c.HTML(http.StatusOK, "log&Signup.html", gin.H{})
	})

	// crée une route /get pour afficher la page
	r.GET("/dashboard-login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "log&Signup.html", gin.H{})
	})

	// permet de vérifier les informations du méthod POST
	r.POST("/login", func(ctx *gin.Context) {
		loginService := service.StaticLoginService(ctx, dbConnector)
		var jwtService service.JWTService = service.JWTAuthService()
		var loginController controller.LoginController = controller.LoginHandler(loginService, jwtService)
		token := loginController.Login(ctx)
		if token != "" {
			fmt.Println("Token found")
			ctx.SetCookie("token", token, 3600, "/", "", false, false)
			ctx.Redirect(http.StatusFound, "/index")
		} else {
			fmt.Println("Token not found")
			ctx.Redirect(http.StatusFound, "/dashboard-login")
		}
	})

	// Middleware pour vérifier le token à chaque requête
	r.Use(func(c *gin.Context) {
		token, err := c.Cookie("token")
		if err != nil || token == "" {
			// Rediriger vers la page de connexion si le token est invalide ou non présent
			c.Redirect(http.StatusFound, "/authentificationfield")
			c.Abort()
			return
		}
		c.Next()
	})
	r.GET("/index", func(c *gin.Context) {

		posts, err := database.GetAllPosts(dbConnector)
		if err != nil {
			// Gérer l'erreur de récupération des posts
			fmt.Printf("Erreur de récupération des posts : %s", err)
			// Par exemple, renvoyer une erreur ou effectuer une autre action appropriée
		}
		c.HTML(http.StatusOK, "index.html", gin.H{
			"posts": posts,
		})
	})

	r.GET("/logout", func(c *gin.Context) {
		c.HTML(http.StatusOK, "log&Signup.html", gin.H{})
	})

	r.GET("/create-post", func(c *gin.Context) {
		c.HTML(http.StatusOK, "create-post.html", gin.H{})
	})

	r.POST("/sendPost", func(c *gin.Context) {
		service.CreatePost(c, dbConnector)
		c.Redirect(http.StatusFound, "/index")
	})

	r.POST("/sendComment", func(c *gin.Context) {
		service.addComment(c, dbConnector)
		c.Redirect(http.StatusFound, "/index")
	})
	r.Run(":8089")
}
