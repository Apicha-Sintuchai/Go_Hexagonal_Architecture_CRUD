package adapters

import (
	"fmt"
	"go/clean/core"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type HttpTodoHandler struct {
	service core.TodolistService
}

func NewHttpTodoHandler(service core.TodolistService) *HttpTodoHandler {
	return &HttpTodoHandler{service: service}
}

func (h *HttpTodoHandler) CreateTodo(c *gin.Context) {
	var todo core.Todolist
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := h.service.CreateTodo(todo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create todo"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"Create": "Sucessful"})
}

func (h *HttpTodoHandler) GetTodo(c *gin.Context) {
	todos, err := h.service.GetTodo()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get todos"})
		return
	}
	c.JSON(http.StatusOK, todos)
}

func (h *HttpTodoHandler) GetID(c *gin.Context) {
	idparam := c.Param("id")
	id, err := strconv.Atoi(idparam)
	if err != nil {
		panic(err)
	}
	todo, err := h.service.GetID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to getbyid todo"})
		return
	}
	c.JSON(http.StatusOK, todo)
}

func (h *HttpTodoHandler) DeleteID(c *gin.Context) {
	idparam := c.Param("id")
	id, err := strconv.Atoi(idparam)
	if err != nil {
		panic(err)
	}
	todo, err := h.service.DeleteID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to Delete todo"})
		return
	}
	fmt.Println(todo)
	c.JSON(http.StatusOK, gin.H{"delete": "sucessful"})
}

func (h *HttpTodoHandler) UpdateTodo(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	var updatedTodo core.Todolist
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	if err := c.ShouldBindJSON(&updatedTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	todo, err := h.service.Update(id, updatedTodo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update todo"})
		return
	}
	fmt.Println(todo)
	c.JSON(http.StatusOK, gin.H{"message": "Update successful"})
}
