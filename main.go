package main

import (
	"html/template"
	"log"
	"net/http"
	"practica7/database"
	"practica7/usuario"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	database.ConnectDB()
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/api/users", usuario.UsersHandler)
	http.HandleFunc("/api/users/", usuario.UserHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Println("Servidor corriendo en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	users, err := usuario.GetAllUsers()
	if err != nil {
		http.Error(w, "Error al obtener usuarios", http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"title":       "Practica 7",
		"total_users": len(users),
		"users":       users,
	}

	templates.ExecuteTemplate(w, "index.html", data)
}
