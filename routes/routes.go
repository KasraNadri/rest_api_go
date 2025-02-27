package routes

import "github.com/gin-gonic/gin"

func RegisterEvents(server *gin.Engine) {
	server.GET("/events", getEvents)       // --- GET
	server.GET("/events/:id", getEvent)    // --- GET(BY ID)
	server.POST("/events", createEvent)    // --- POST
	server.PUT("/events/:id", updateEvent) // --- PUT
}
