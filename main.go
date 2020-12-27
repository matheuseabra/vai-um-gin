package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// User struct
type User struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func createUserHandler(c *gin.Context) {
	var user User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Not a JSON request body"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

// BootstrapServer boostraps the gin router
func BootstrapServer() *gin.Engine {
	router := gin.Default()

	router.GET("/api/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "The API is up",
		})
	})

	router.POST("/api/session", createUserHandler)

	return router
}

func main() {

	server := BootstrapServer()

	server.Run(":9000")
}
