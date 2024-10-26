package usuario

import (
	"encoding/json"
	"net/http"
	"practica7/database"
	"strconv"
)

func GetAllUsers() ([]User, error) {
	rows, err := database.DB.Query("SELECT id, name, email FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []User{}
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		users, err := GetAllUsers()
		if err != nil {
			http.Error(w, "Error al obtener usuarios", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(users)
	case "POST":
		var user User
		json.NewDecoder(r.Body).Decode(&user)
		_, err := database.DB.Exec("INSERT INTO users (name, email) VALUES (?, ?)", user.Name, user.Email)
		if err != nil {
			http.Error(w, "Error al crear usuario", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
	}
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Path[len("/api/users/"):])
	if r.Method == "DELETE" {
		_, err := database.DB.Exec("DELETE FROM users WHERE id = ?", id)
		if err != nil {
			http.Error(w, "Error al eliminar usuario", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}
