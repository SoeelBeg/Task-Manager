package controller

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"task-manager/models"
)

var tmpl *template.Template

func init() {
	var err error
	// Update the path to the template
	tmpl, err = template.ParseFiles("../../templates/tasks.html")
	if err != nil {
		log.Fatalf("Error loading template: %v", err)
	}
}

func SetDatabase(database *sql.DB) {
	models.SetDatabase(database)
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	tasks, err := models.GetTasks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.ExecuteTemplate(w, "tasks.html", tasks)
}

func AddTask(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		name := r.FormValue("task")
		err := models.AddTask(name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		idStr := r.FormValue("task_id")
		id, _ := strconv.Atoi(idStr)
		err := models.DeleteTask(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
