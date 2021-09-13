package services

import (
	"net.codefiction/go-tutorial/database"
	"net.codefiction/go-tutorial/dto"
	"net.codefiction/go-tutorial/models"
)

func CreateTodo(todo dto.TodoInput) {
	db := database.ConnectDataBase()

	insertTodo, err := db.Prepare("insert into todos(description) values ($1)")
	if err != nil {
		panic(err.Error())
	}

	_, err = insertTodo.Exec(todo.Description)
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
}

func DeleteTodo(id string) {
	db := database.ConnectDataBase()

	deleteTodo, err := db.Prepare("delete from todos where id=$1")
	if err != nil {
		panic(err.Error())
	}

	_, err = deleteTodo.Exec(id)
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
}

func ListTodos() *[]models.Todo {
	db := database.ConnectDataBase()

	selectTodos, err := db.Query("SELECT * FROM todos")
	if err != nil {
		panic(err.Error())
	}

	var todos []models.Todo
	var todo models.Todo

	for selectTodos.Next() {
		var id int
		var description string

		err = selectTodos.Scan(&id, &description)
		if err != nil {
			panic(err.Error())
		}

		todo.Id = id
		todo.Description = description

		todos = append(todos, todo)
	}

	defer db.Close()

	return &todos
}
