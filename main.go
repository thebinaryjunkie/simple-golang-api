package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)
type user struct {
	Id			string	`json:"id"`
	Name		string	`json:"name"`
	Age			int		`json:"age"`
	Member		bool	`json:"member"`
}

var users = []user {
	{Id: "1", Name: "Joe", Age: 29, Member: false},
	{Id: "2", Name: "Rick", Age: 55, Member: false},
	{Id: "3", Name: "Patty", Age: 21, Member: true},
}

func getUsers(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, users)
}

func main() {

	router := gin.Default()
	router.GET("/users", getUsers)
	router.Run("localhost:8080")
}