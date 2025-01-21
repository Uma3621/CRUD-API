package main

import (
	"example.com/crud-api/config"
	"example.com/crud-api/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()
	router := gin.Default()

	router.POST("/user", handlers.CreateUser)
	router.GET("/users", handlers.GetUsers)
	router.PUT("/users/:id", handlers.UpdateUser)
	router.DELETE("/users/:id", handlers.DeleteUser)
	router.Run(":8080")
}

/*type LoginRequest struct {
	Name     string `json:"name" binding:"required"` //json key
	Password string `json:"password" binding:"required"`
}

func main() {
	router := gin.Default() //created default gin instance

	router.POST("/login", func(c *gin.Context) { //define post endurl and
		//handler function should used to help http request&responses

		var loginrequest LoginRequest //create loginrequest var to store json payload

		if err := c.ShouldBindJSON(&loginrequest); err != nil { //read json and binds it to struct
			c.JSON(http.StatusBadRequest, gin.H{
				"Error": "Invalid request",
			})
			return

		}
		//validating using dummy values
		if loginrequest.Name == "admin" && loginrequest.Password == "password123" {
			c.JSON(http.StatusOK, gin.H{
				"success": "login is successful",
				"token":   "acess token",
			})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "inavlid username or password",
			})
		}

	})
	router.Run(":8080")
}*/
