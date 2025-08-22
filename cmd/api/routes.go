package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (app *application) routes() http.Handler {
	g := gin.Default()

	v1 := g.Group("/v1")
	{
		v1.GET("/users", app.getAllEvents)
		v1.GET("/users/:id", app.getByIdEvent)
		v1.POST("/users", app.createEvent)
		v1.PUT("/users/:id", app.updateEvent)
		v1.DELETE("/users/:id", app.deleteEvent)

		v1.POST("/auth/register", app.registerUser)
		// v1.POST("/auth/login", app.loginUser)
	}

	return g
}
