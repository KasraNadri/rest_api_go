package main

import (
	"rest_api/db"
	"rest_api/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	// -- INITIALIZES THE DATABASE -- \\
	db.InitDB()

	// -- INITIALLIZES A SERVER -- \\
	server := gin.Default()

	// -- REGISTERS 'EVENT' ROUTES -- \\
	routes.RegisterEvents(server)

	// -- REGISTERS 'USER' ROUTES -- \\
	routes.RegisterUsers(server)

	// -- STARTS LISTENING AT PORT 'LOCALHOST:8080' -- \\
	server.Run(":8080")

}
