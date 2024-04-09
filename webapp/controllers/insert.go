package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"webapp/structures"
)

func Insert(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var user structures.User

		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, "error decoding json", http.StatusInternalServerError)
			return
		}
		// fmt.Fprintf(w, "the struct is %v", user)
		qry := `INSERT INTO users (username,password,role) VALUES ($1,$2,$3) RETURNING id`
		err = db.QueryRow(qry, user.Username, user.Password, user.Role).Scan(&user.ID)
		if err != nil {
			if err == sql.ErrNoRows {
				http.Error(w, "invalid user", http.StatusBadRequest)
			} else {
				http.Error(w, "error querying data", http.StatusInternalServerError)
			}
			return
		}
		fmt.Fprintf(w, "the inserted user %v ", user)

	}
}
