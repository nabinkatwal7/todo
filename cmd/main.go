package main

import (
	"fmt"
	"log"
	"todo-app/server"
	"todo-app/server/database"
)

func main() {
	database.InitDB()

	r := server.SetupRouter()

	fmt.Println("Server is running on port 8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}