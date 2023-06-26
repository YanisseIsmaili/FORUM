package main

import (
	"Forum/Backend/controller"
	db "Forum/Backend/database"

	"Forum/Backend/service"
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
	db.CreateDB(dbConnector)

	r := gin.Default()
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Printf("endpoint %v %v %v %v\n", httpMethod, absolutePath, handlerName, nuHandlers)
	}
	r.StaticFile("/styles.css", "./Frontend/styles.css") // permet de link le fichier css
	r.LoadHTMLGlob("Frontend/*")                         // permet d'allez chercher

	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "log&Signup.html", gin.H{})
	})

	r.POST("/login", func(ctx *gin.Context) {
		loginService := service.StaticLoginService(ctx, dbConnector)
		var jwtService service.JWTService = service.JWTAuthService()
		var loginController controller.LoginController = controller.LoginHandler(loginService, jwtService)
		token := loginController.Login(ctx)
		if token != "" {
			fmt.Println("Token found")
			ctx.SetCookie("token", token, 3600, "/", "", false, true)
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

	r.GET("/index", func(c *gin.Context) {
		// Code pour gérer l'accès à la page /index lorsque le token est valide
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	
	r.Run(":8089")
}
