package main

import (
	"github.com/gin-gonic/gin"
	"net.codefiction/go-tutorial/controllers"
)

func main() {
	router := gin.Default()

	router.GET("/todos", controllers.ListTodos)
	router.POST("/todos", controllers.CreateTodo)
	router.DELETE("/todos/:id", controllers.DeleteTodo)

	router.Run()
}
