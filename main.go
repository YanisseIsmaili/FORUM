package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

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

		// Récupérer les commentaires pour chaque post
		for i, post := range posts {
			comments, err := service.GetCommentsByPostID(post.ID, dbConnector)
			if err != nil {
				// Gérer l'erreur de récupération des commentaires
				fmt.Printf("Erreur de récupération des commentaires : %s", err)
				// Par exemple, renvoyer une erreur ou effectuer une autre action appropriée
			}
			// Ajouter les commentaires au post correspondant
			posts[i].CommentsUser = comments
		}

		service.ShowPostModal(c, dbConnector) // Appel de la fonction ShowPostModal
		c.HTML(http.StatusOK, "index.html", gin.H{
			"posts": posts,
		})
	})

	r.GET("/post-modal", func(c *gin.Context) {
		postID := c.Query("postID") // Récupérer la valeur de postID depuis les paramètres de requête
		service.ShowPostModal(c, dbConnector)
		c.HTML(http.StatusOK, "post-modal.html", gin.H{
			"postID": postID,
		})

	})

	r.GET("/comments/:postID", func(c *gin.Context) {
		postIDStr := c.Param("postID")
		postID, err := strconv.ParseUint(postIDStr, 10, 64)
		if err != nil {
			// Gérer l'erreur de conversion de l'ID du post
			fmt.Printf("Erreur de conversion de l'ID du post : %s", err)
			// Par exemple, renvoyer une erreur ou effectuer une autre action appropriée
			return
		}

		comments, err := service.GetCommentsByPostID(uint(postID), dbConnector)
		if err != nil {
			// Gérer l'erreur de récupération des commentaires
			fmt.Printf("Erreur de récupération des commentaires : %s", err)
			// Par exemple, renvoyer une erreur ou effectuer une autre action appropriée
		}

		c.HTML(http.StatusOK, "comments.html", gin.H{
			"comments": comments,
		})
	})

	r.GET("/logout", func(c *gin.Context) {
		c.HTML(http.StatusOK, "log&Signup.html", gin.H{})
	})

	r.GET("/create-post", func(c *gin.Context) {
		c.HTML(http.StatusOK, "create-post.html", gin.H{})
	})

	r.POST("/comments", func(c *gin.Context) {
		// Lire le corps de la requête
		body, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur de lecture du corps de la requête"})
			return
		}

		// Afficher le contenu du corps de la requête
		c.JSON(http.StatusOK, gin.H{"content": string(body)})
	})

	// r.POST("/sendComment", func(c *gin.Context) {
	// 	// Récupérer les données du formulaire
	// 	comment := c.PostForm("comment")
	// 	postIDStr := c.PostForm("postID")

	// 	// Convertir postIDStr en uint
	// 	postID, err := strconv.ParseUint(postIDStr, 10, 64)
	// 	if err != nil {
	// 		// Gérer l'erreur de conversion
	// 		// Par exemple, renvoyer une réponse d'erreur
	// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid postID"})
	// 		return
	// 	}

	// 	// Appeler la fonction SendComment du service avec les valeurs appropriées
	// 	err = service.SendComment(uint(postID), comment, dbConnector)
	// 	if err != nil {
	// 		// Gérer l'erreur de l'envoi du commentaire
	// 		// Par exemple, renvoyer une réponse d'erreur
	// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send comment"})
	// 		return
	// 	}

	// 	// Répondre avec une réponse réussie
	// 	c.JSON(http.StatusOK, gin.H{"message": "Comment sent successfully"})
	// })
	r.Run(":8089")

}
