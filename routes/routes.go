package routes

import "github.com/gin-gonic/gin"

// ========== EVENT ROUTES REGISTERER ========== \\
func RegisterEvents(server *gin.Engine) {
	server.GET("/events", getEvents)          // --- GET
	server.GET("/events/:id", getEvent)       // --- GET(BY ID)
	server.POST("/events", createEvent)       // --- POST
	server.PUT("/events/:id", updateEvent)    // --- PUT(BY ID)
	server.DELETE("/events/:id", deleteEvent) // --- DELETE(BY ID)
}

// ========== USER ROUTES REGISTERER ========== \\
func RegisterUsers(server *gin.Engine) {
	server.POST("/signup", signUp) // --- SIGNS UP USERS
	server.POST("/login", logIn)   // --- LOGS IN USERS
}
