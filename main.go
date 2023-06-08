package main

import ("github.com/gin-gonic/gin"
		"net/http"
		"log"
)

func main() {
	// route par default : 
	r := gin.Default()
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Printf("endpoint %v %v %v %v\n", httpMethod, absolutePath, handlerName, nuHandlers)
	}
	r.StaticFile("/styles.css", "./Frontend/styles.css")
	r.LoadHTMLGlob("Frontend/*")
	r.GET("/loging", func(c *gin.Context) { // récupére le formulaire : 
		c.HTML(http.StatusOK, "log&Signup.html", gin.H{
			"title": "Main website",
		})
	})
	r.Run()
}