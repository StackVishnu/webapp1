package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"webapp/structures"
)

func ViewAll(db *sql.DB) http.HandlerFunc {
	var mobiles []structures.Mobile
	return func(w http.ResponseWriter, r *http.Request) {
		qry := `SELECT id,name,specs,price,image_url FROM products`
		rows, err := db.Query(qry)
		if err != nil {
			http.Error(w, "error in query", http.StatusInternalServerError)
			return
		}
		for rows.Next() {
			var mob structures.Mobile
			err := rows.Scan(&mob.ID, &mob.Name, &mob.Specs, &mob.Price, &mob.Ipath)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			mobiles = append(mobiles, mob)
		}
		// fmt.Fprintf(w, "the slice has %v", mobiles)
		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(mobiles)
		if err != nil {
			http.Error(w, "error encoding json", http.StatusInternalServerError)
			return
		}
		mobiles = nil
	}
}
