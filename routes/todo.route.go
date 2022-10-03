package routes

import (
	"trandung/server/controllers"

	"github.com/gin-gonic/gin"
)

func AddTodoRoutes(rg *gin.RouterGroup) {
	todos := rg.Group("/todos")
	todos.GET("/", controllers.GetTodos)
	todos.POST("/", controllers.AddTodo)
	todos.GET("/:id", controllers.GetTodoById)
	todos.DELETE("/:id", controllers.DeleteTodoById)
	todos.PUT("/:id", controllers.UpdateById)
}
