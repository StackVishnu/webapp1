package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"webapp/structures"
)

func Search(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Decode the JSON request body into a struct
		var searchKey struct {
			Key string `json:"key"`
		}
		err := json.NewDecoder(r.Body).Decode(&searchKey)
		if err != nil {
			http.Error(w, "error decoding JSON: "+err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Printf("Received search key: %s\n", searchKey.Key)

		// Prepare the query
		qry := `SELECT * FROM products WHERE name LIKE '%' || $1 || '%'`

		// Execute the query
		rows, err := db.Query(qry, searchKey.Key)
		if err != nil {
			http.Error(w, "error executing query: "+err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		// Iterate over the rows and populate the struct

		var res []structures.Mobile
		for rows.Next() {
			var mob structures.Mobile
			err := rows.Scan(&mob.ID, &mob.Name, &mob.Specs, &mob.Price, &mob.Ipath)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			res = append(res, mob)
		}
		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(res)
		if err != nil {
			http.Error(w, "error encoding json", http.StatusInternalServerError)
			return
		}
	}
}
