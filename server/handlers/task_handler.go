package handlers

import (
	"fmt"
	"net/http"
	"todo-app/server/database"
	"todo-app/server/models"

	"github.com/gin-gonic/gin"
)

func CreateTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBind(&task); err != nil {
		c.String(http.StatusBadRequest, "Invalid request")
		return
	}

	database.DB.Create(&task)

	c.String(http.StatusOK, renderTaskHTML(task))
}

func GetTasks(c *gin.Context) {
	var tasks []models.Task
	database.DB.Find(&tasks)

	var taskListHTML string
	for _, task := range tasks {
		taskListHTML += renderTaskHTML(task)
	}

	c.String(http.StatusOK, taskListHTML)
}

func MarkTaskComplete(c *gin.Context) {
	id := c.Param("id")
	var task models.Task

	if err := database.DB.First(&task, id).Error; err != nil {
		c.String(http.StatusNotFound, "Task not found")
		return
	}

	task.Completed = true
	database.DB.Save(&task)

	c.String(http.StatusOK, renderTaskHTML(task))
}

func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	var task models.Task

	if err := database.DB.First(&task, id).Error; err != nil {
		c.String(http.StatusNotFound, "Task not found")
		return
	}

	database.DB.Delete(&task)

	c.String(http.StatusOK, "")
}

func renderTaskHTML(task models.Task) string {
	completedClass := ""
	if task.Completed {
		completedClass = "line-through text-gray-500"
	}

	return fmt.Sprintf(`
		<li id="task-%d" class="flex justify-between bg-gray-100 p-2 rounded">
			<span class="%s">%s</span>
			<div>
				<button hx-put="/tasks/%d/complete" hx-target="#task-%d" hx-swap="outerHTML" class="text-green-600">✔</button>
				<button hx-delete="/tasks/%d" hx-target="#task-%d" hx-swap="outerHTML" class="text-red-600">✖</button>
			</div>
		</li>`, task.ID, completedClass, task.Title, task.ID, task.ID, task.ID, task.ID)
}
