package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (app *application) routes() http.Handler {
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
