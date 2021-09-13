package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"net.codefiction/go-tutorial/dto"
	"net.codefiction/go-tutorial/services"
)

func CreateTodo(c *gin.Context) {
	var todoInput dto.TodoInput

	if err := c.ShouldBindJSON(&todoInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	services.CreateTodo(todoInput)

	c.JSON(http.StatusOK, gin.H{"message": "created"})
}

func DeleteTodo(c *gin.Context) {
	id := c.Param("id")

	services.DeleteTodo(id)

	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}

func ListTodos(c *gin.Context) {

	todos := services.ListTodos()

	c.JSON(http.StatusOK, todos)
}
