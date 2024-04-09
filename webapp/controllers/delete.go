package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"webapp/structures"
)

func Delete(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user structures.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, "error decoding JSON", http.StatusInternalServerError)
			return
		}

		// Retrieve username before deletion
		var username string
		qry := `SELECT username FROM users WHERE id = $1`
		err = db.QueryRow(qry, user.ID).Scan(&username)
		if err != nil {
			if err == sql.ErrNoRows {
				http.Error(w, "no such user", http.StatusBadRequest)
			} else {
				http.Error(w, "error querying data", http.StatusInternalServerError)
			}
			return
		}

		// Delete user
		qry = `DELETE FROM users WHERE id = $1`
		_, err = db.Exec(qry, user.ID)
		if err != nil {
			http.Error(w, "error deleting user", http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "Deleted user with username: %s", username)
	}
}
