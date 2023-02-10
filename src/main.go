package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"errors"
	"go-project/users"
)

var Users = users.GetAllUsers()

func getUsers (context *gin.Context) {
	context.IndentedJSON(http.StatusOK, Users)
}

func addUser (context *gin.Context) {
	var newUser users.User

	if err := context.BindJSON(&newUser); err != nil {
		return
	}

	Users = append(Users, newUser)

	context.IndentedJSON(http.StatusCreated, newUser)
}

func updateUser (context *gin.Context) {
	id := context.Param("id")

	user, err := getUserById(id)

	var newUser users.User

	if err != nil{
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "User not found"})
	} else {
		if err := context.BindJSON(&newUser); err != nil {
			return
		}

		user.Id = newUser.Id
		user.Name = newUser.Name
		user.Tel = newUser.Tel
		user.Email = newUser.Email
		user.Paswsord = newUser.Paswsord
	}
}

func getUserById (id string) (*users.User, error) {
	for i, U:= range Users {
		if U.Id == id {
			return &Users[i], nil
		}
	}

	return nil, errors.New("User not found")
}

func getUser (context *gin.Context) {
	id := context.Param("id")

	user, err := getUserById(id)

	if err != nil{
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "User not found"})
	} else {
		context.IndentedJSON(http.StatusOK, user)
	}
}

func main() {
	router := gin.Default()
	router.GET("/users", getUsers)
	router.GET("/users/:id", getUser)
	router.PATCH("/users/:id", updateUser)
	router.POST("/users", addUser)
	router.Run("localhost:9090")
}