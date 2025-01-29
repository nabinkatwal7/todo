package main

import (
	"fmt"
	"log"
	"net/http"
	"todo-app/server"

	"github.com/gin-gonic/gin"
)

func main() {
	server.InitDB()

	r := gin.Default()

	r.GET("/", func (c *gin.Context) {
		c.String(http.StatusOK, "Todo app is running")
	})

	fmt.Println("Server is running on port 8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}