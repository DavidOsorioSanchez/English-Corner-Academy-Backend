package main

import (
	"net/http"

	// "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (app *Application) routes() http.Handler {
	g := gin.Default()

	mainGroup := g.Group("/v1")
	{
		// Events
		mainGroup.GET("/events", app.getAllEvents)
		mainGroup.GET("/events/:id", app.getByIdEvent)
		mainGroup.GET("/events/:id/attendees", app.getEventAttendeeForEvent)
		mainGroup.GET("/attendees/:id/events", app.getEventsByAttendee)

		mainGroup.POST("/auth/register", app.registerUser)
		mainGroup.POST("/auth/login", app.loginUser)
		// mainGroup.POST("/auth/register/google", app.registerGoogle)
	}

	authGroup := mainGroup.Group("/")
	authGroup.Use(app.authMiddleware())

	// necesito acordarme de descomentar esto si quiero probar desde el front
	// ademas de que me gusta la idea de usar cors
	// ya que me agrega una seguridad contra los metodos que puede usar y desde donde se puede acceder

	// , cors.New(cors.Config{
	// // establecer las configuraciones de cors, necesito acordarme cambiar el origen cuando suba el front a produccion
	// 	AllowOrigins: []string{"http://localhost:5173"},
	// 	AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
	// 	// AllowHeaders:     []string{"Accept", "Authorization", "Content-Type"},
	// 	AllowCredentials: true, //cookies permitidas
	// })
	// )
	{
		authGroup.POST("/events", app.createEvent)
		authGroup.PUT("/events/:id", app.updateEvent)
		authGroup.DELETE("/events/:id", app.deleteEvent)
		authGroup.POST("/events/:id/attendees/:userId", app.addAttendeeToEvent)
		authGroup.DELETE("/events/:id/attendees/:userId", app.deleteAttendeeFromEvent)
	}

	g.GET("/swagger/*any", func(c *gin.Context) {
		if c.Request.RequestURI == "/swagger/" {
			c.Redirect(302, "/swagger/index.html")
		}
		ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL("http://localhost:8080/swagger/doc.json"))(c)
	})

	return g
}
