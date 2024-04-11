package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"webapp/structures"

	"github.com/gorilla/mux"
)

func ViewEach(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var mob structures.Mobile
		params := mux.Vars(r)
		key, _ := strconv.Atoi(params["id"])
		qry := `SELECT * FROM products WHERE id = $1`
		rows := db.QueryRow(qry, key)
		rows.Scan(&mob.ID, &mob.Name, &mob.Specs, &mob.Price, &mob.Ipath)
		var arr []structures.Mobile
		arr = append(arr, mob)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(&arr)

	}
}
