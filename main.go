package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"errors"
)

//models
type user struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Member bool   `json:"member"`
}

var users = []user{
	{Id: "1", Name: "Joe", Age: 29, Member: false},
	{Id: "2", Name: "Rick", Age: 55, Member: false},
	{Id: "3", Name: "Patty", Age: 21, Member: true},
}

func main() {

	router := gin.Default()
	router.GET("/users", getUsers)
	router.GET("/users/:id", getUser)
	router.POST("/sign-up/:id", signUpUser)
	router.Run("localhost:8080")
}

//handlers
func getUsers(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, users)
}

func getUser(context *gin.Context) {
	id := context.Param("id")
	user, err := getUserById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
	}

	context.IndentedJSON(http.StatusOK, user)
}

func signUpUser(context *gin.Context) {
	id := context.Param("id")
	user, err := getUserById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
	}

	user.Member = true

	context.IndentedJSON(http.StatusOK, user)
}

//util
func getUserById(id string) (*user, error) {

	for i, u := range users {
		if u.Id == id {
			return &users[i], nil
		}
	}
	return nil, errors.New("user does not exist")
}
