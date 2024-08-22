package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"task-manager/controller"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Set working directory to the project root
	os.Chdir("../../")

	// Database connection
	db, err := sql.Open("mysql", "admin:soeel123@tcp(127.0.0.1:3306)/taskmanager")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Pass the database connection to the handlers
	controller.SetDatabase(db)

	http.HandleFunc("/", controller.HomePage)
	http.HandleFunc("/add-task", controller.AddTask)
	http.HandleFunc("/delete-task", controller.DeleteTask)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
