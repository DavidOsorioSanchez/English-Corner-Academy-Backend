package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
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

	return g
}
