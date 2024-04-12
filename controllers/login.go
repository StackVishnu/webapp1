package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"webapp/structures"
)

// login page is at http://127.0.0.1:5500/frontend/ (if live @ 5500)
func Login(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user structures.User
		err := r.ParseMultipartForm(10 << 20)
		if err != nil {
			http.Error(w, "Error parsing form data", http.StatusBadRequest)
			return
		}

		username := r.Form.Get("username")
		password := r.Form.Get("password")

		if password == "" || username == "" {
			http.Error(w, "username and password required", http.StatusBadRequest)
			return
		}
		// w.Header().Set("Content-Type", "applicaton/json")
		query := `SELECT role FROM users WHERE  username = $1 AND password =$2`
		err = db.QueryRow(query, username, password).Scan(&user.Role)
		if err != nil {
			if err != sql.ErrNoRows {
				http.Error(w, "INVALID USER", http.StatusUnauthorized)
			} else {
				http.Error(w, "Database error", http.StatusInternalServerError)
			}
			return
		}
		response := map[string]string{
			"role": user.Role,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)

	}
}
