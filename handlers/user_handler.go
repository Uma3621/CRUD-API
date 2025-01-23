package handlers

import (
	"context"
	"log"
	"net/http"
	"time"

	"example.com/crud-api/Configurations"
	"example.com/crud-api/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//var config Configurations.Config

// CreateUser handles creating a new user.
func CreateUser(c *gin.Context, config Configurations.Config) {
	var user models.User

	// Binds the JSON body to the `user` struct.
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Access the "users" collection in MongoDB.
	collection := Configurations.GetCollection(config.Database.Name, "users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Insert the user into the collection.
	result, err := collection.InsertOne(ctx, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	// Return the inserted ID in the response.
	c.JSON(http.StatusCreated, gin.H{"id": result.InsertedID})
}

// GetUsers retrieves all users.
func GetUsers(c *gin.Context, config Configurations.Config) {
	collection := Configurations.GetCollection(config.Database.Name, "users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{}) // Fetch all documents using an empty filter.
	if err != nil {
		log.Printf("Error fetching users from DB: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}
	defer cursor.Close(ctx)

	var users []models.User
	if err := cursor.All(ctx, &users); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse users"})
		return
	}

	c.JSON(http.StatusOK, users) // Returns the list of users in JSON format.
}

// UpdateUser updates a user by ID.
func UpdateUser(c *gin.Context, config Configurations.Config) {
	id := c.Param("id") // Retrieves the user ID from the URL path parameter.
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	objectID, _ := primitive.ObjectIDFromHex(id) // Converts the ID string to MongoDB's ObjectID type.
	collection := Configurations.GetCollection(config.Database.Name, "users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := collection.UpdateOne(ctx, bson.M{"_id": objectID}, bson.M{"$set": user})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

// DeleteUser deletes a user by ID.
func DeleteUser(c *gin.Context, config Configurations.Config) {
	id := c.Param("id")
	objectID, _ := primitive.ObjectIDFromHex(id)
	collection := Configurations.GetCollection(config.Database.Name, "users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := collection.DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
