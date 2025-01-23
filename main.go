package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"example.com/crud-api/Configurations"
	"example.com/crud-api/handlers"
	"github.com/gin-gonic/gin"
)

// loading the config file
var config Configurations.Config

func LoadConfig(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	err := LoadConfig("config.json") // Load the config file
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}
	// Connect to the database using config values
	Configurations.ConnectDatabase(config)
	router := gin.Default()

	router.POST("/user", func(c *gin.Context) {
		handlers.CreateUser(c, config) // Pass the required `*gin.Context` and `config`.
	})
	router.GET("/users", func(c *gin.Context) {
		handlers.GetUsers(c, config) // Pass the required `*gin.Context` and `config`.
	})
	router.PUT("/users/:id", func(c *gin.Context) {
		handlers.UpdateUser(c, config) // Pass the required `*gin.Context` and `config`.
	})
	router.DELETE("/users/:id", func(c *gin.Context) {
		handlers.DeleteUser(c, config) // Pass the required `*gin.Context` and `config`.
	})
	// Start the server using the port from the config
	ServerAddress := fmt.Sprintf(":%d", config.Server.Port)
	router.Run(ServerAddress)
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
