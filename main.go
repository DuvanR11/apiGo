package main

import (
	"net/http"
	"os"
	"github.com/DuvanR11/api/db"
	"github.com/DuvanR11/api/models"
	"github.com/DuvanR11/api/routes"
	"github.com/gorilla/mux"
)

func main() {
	//Conección a la base de datos
	db.DBConnection()

	//Migraciones para crear las tablas
	db.DB.AutoMigrate(models.Task{})
	db.DB.AutoMigrate(models.User{})

	//URL
	r := mux.NewRouter()
	//Principal
	r.HandleFunc("/", routes.Home)

	//CRUD USUARIO
	r.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	r.HandleFunc("/users/{id}", routes.GetUserHandler).Methods("GET")
	r.HandleFunc("/users", routes.PostUsersHandler).Methods("POST")
	r.HandleFunc("/users/{id}", routes.DeleteUsersHandler).Methods("DELETE")
	r.HandleFunc("/users/{id}", routes.UpdatetUsersHandler).Methods("PUT")

	//CRUD TAREAS
	r.HandleFunc("/tasks", routes.GetTasksHandler).Methods("GET")
	r.HandleFunc("/tasks/{id}", routes.GetTaskHandler).Methods("GET")
	r.HandleFunc("/tasks", routes.PostTaskHandler).Methods("POST")
	r.HandleFunc("/tasks/{id}", routes.DeleteTaskHandler).Methods("DELETE")
	r.HandleFunc("/tasks/{id}", routes.UpdateTaskHandler).Methods("PUT")

	//Puerto
	// http.ListenAndServe(":3000", r)
	http.ListenAndServe(":" + os.Getenv("PORT"), r)
}
