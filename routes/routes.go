package routes

import "github.com/gin-gonic/gin"

func RegisterEvents(server *gin.Engine) {
	server.GET("/events", getEvents)          // --- GET
	server.GET("/events/:id", getEvent)       // --- GET(BY ID)
	server.POST("/events", createEvent)       // --- POST
	server.PUT("/events/:id", updateEvent)    // --- PUT(BY ID)
	server.DELETE("/events/:id", deleteEvent) // --- DELETE(BY ID)
}

func RegisterUsers(server *gin.Engine) {
	server.POST("/signup") // --- SIGNS UP USERS
}
