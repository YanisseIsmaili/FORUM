package main

import (
	 Forum "Forum/Backend"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// route par default :
	r := gin.Default()
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Printf("endpoint %v %v %v %v\n", httpMethod, absolutePath, handlerName, nuHandlers)
	}
	r.StaticFile("/styles.css", "./Frontend/styles.css") // permet de link le fichier css
	r.LoadHTMLGlob("Frontend/*")                         // permet d'allez chercher
	r.GET("/loging", func(c *gin.Context) {
		c.HTML(http.StatusOK, "log&Signup.html", gin.H{
			"title": "Main website",
		})
	})
	r.POST("/submit", func(c *gin.Context) {
		username := c.PostForm("username")
		email := c.PostForm("email")
		password := c.PostForm("password")

		c.String(200, "Hello %s", username, email, password)
	})

	r.Run()
Forum.createDB()
	


}
