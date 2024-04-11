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
		var mob structures.Mobile
		err := json.NewDecoder(r.Body).Decode(&mob)
		fmt.Printf("%v", mob)
		if err != nil {
			http.Error(w, "error decoding JSON", http.StatusInternalServerError)
			return
		}

		qry := `SELECT name FROM products WHERE id = $1`
		err = db.QueryRow(qry, mob.ID).Scan(&mob.Name)
		if err != nil {
			if err == sql.ErrNoRows {
				http.Error(w, "no such user", http.StatusBadRequest)
			} else {
				http.Error(w, "error querying data", http.StatusInternalServerError)
			}
			return
		}

		// Delete user
		qry = `DELETE FROM products WHERE id = $1`
		_, err = db.Exec(qry, mob.ID)
		if err != nil {
			http.Error(w, "error deleting user", http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(mob)

	}
}
