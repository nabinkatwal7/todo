package server

import (
	"todo-app/server/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()


	r.GET("/", func(c *gin.Context) {
		c.File("./templates/index.html")
	})

	r.POST("/tasks", handlers.CreateTask)
	r.GET("/tasks", handlers.GetTasks)
	r.PUT("/tasks/:id/complete", handlers.MarkTaskComplete)
	r.DELETE("/tasks/:id", handlers.DeleteTask)

	return r
}
