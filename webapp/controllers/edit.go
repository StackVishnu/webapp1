package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"webapp/structures"
)

func Edit(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var mob structures.Mobile

		idStr := r.FormValue("id")
		id, _ := strconv.Atoi(idStr)
		mob.ID = id

		mob.Name = r.FormValue("name")
		mob.Specs = r.FormValue("specs")

		priceStr := r.FormValue("price")
		price, _ := strconv.Atoi(priceStr)
		mob.Price = price

		qry := `UPDATE  products SET name = $1, specs = $2, price = $3 WHERE id = $4`
		_, err := db.Exec(qry, mob.Name, mob.Specs, mob.Price, mob.ID)
		if err != nil {
			http.Error(w, "error updating mob "+err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "the struct is : %v", mob)
	}
}
