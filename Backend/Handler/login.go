package Forum

import (
	Database "Forum/Backend/Database"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type LoginService interface {
	LoginUser(email string, password string) bool
}

type loginInformation struct {
	emailUser    string
	passwordUser string
}

func StaticLoginService(c *gin.Context) LoginService {
	email := c.PostForm("email")
	password := c.PostForm("password")

	// Création d'un nouvel utilisateur avec les données récupérées
	user := Database.Users{

		Email:    email,
		Password: password,
	}
	return &loginInformation{
		emailUser:    user.Email,
		passwordUser: user.Password,
	}
}
func (info *loginInformation) LoginUser(email string, password string) bool {
	return info.emailUser == email && info.passwordUser == password
}

func Register() {
	// crée la base de donnée :
	dsn := "NerdMythology.db"
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
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "log&Signup.html", gin.H{})
	})
	r.POST("/login", func(c *gin.Context) {

		// username := c.PostForm("username")
		// email := c.PostForm("email")
		// password := c.PostForm("password")

		// // Création d'un nouvel utilisateur avec les données récupérées
		// user := Database.Users{
		// 	Username: username,
		// 	Email:    email,
		// 	Password: password,
		// }

		// Appel à la fonction createDB() pour créer et initialiser l'utilisateur
		Database.CreateDB(db)
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
		c.HTML(http.StatusOK, "log&Signup.html", gin.H{}) // permet d'allez chercher le fichier html

	})

	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	r.Run()
}
