package main

import (
	"log"
	"net/http"
	"note-app/controllers"
	"note-app/middlewares"
	"note-app/models"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(gin.Logger())

	r.Static("/vendor", "./static/vendor")

	r.LoadHTMLGlob("templates/**/*")

	models.ConnectDatabase()
	models.DBMigrate()

	store := memstore.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("notes", store))

	notes := r.Group("/notes")
	{
		notes.GET("/", controllers.NotesIndex)
		notes.GET("/new", controllers.NotesNew)
		notes.POST("/", controllers.NotesCreate)
		notes.GET("/:id", controllers.NotesShow)
		notes.GET("/edit/:id", controllers.NotesEdit)
		notes.POST("/:id", controllers.NotesUpdate)
		notes.DELETE("/:id", controllers.NotesDelete)
	}

	r.GET("/login", controllers.LoginPage)
	r.GET("/signup", controllers.SignupPage)

	r.POST("/signup", controllers.SignUp)
	r.POST("/login", controllers.Login)
	r.POST("/logout", controllers.Logout)

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home/index.html", gin.H{
			"title":     "Notes Appliaction",
			"logged_in": (c.GetUint64("user_id") > 0),
		})
	})

	r.Use(middlewares.AuthenticateUser())

	log.Println("Server Started")
	r.Run(":9090")
}
