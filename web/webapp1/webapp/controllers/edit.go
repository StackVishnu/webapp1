package controllers

import (
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strconv"
	"webapp/structures"
)

func Edit(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var mob structures.Mobile
		err := r.ParseMultipartForm(10 << 20)
		if err != nil {
			http.Error(w, "Error parsing form data", http.StatusBadRequest)
			return
		}
		mob.Name = r.Form.Get("phone_name")
		mob.Price, err = strconv.Atoi(r.Form.Get("price"))
		if err != nil {
			http.Error(w, "Error parsing phone price", http.StatusBadRequest)
			return
		}
		mob.Specs = r.Form.Get("specs")

		file, header, err := r.FormFile("image")
		if err != nil {
			http.Error(w, "Error retrieving image file", http.StatusBadRequest)
			return
		}
		defer file.Close()

		// Save the file to your desired location
		out, err := os.Create("C:/Users/njana/OneDrive/Desktop/web/webapp1/webapp/frontend/assets/" + header.Filename)
		if err != nil {
			http.Error(w, "Error creating file: "+err.Error(), http.StatusBadRequest)
			return
		}
		defer out.Close()

		_, err = io.Copy(out, file)
		if err != nil {
			http.Error(w, "Error copying file", http.StatusBadRequest)
			return
		}

		mob.Ipath = "C:/Users/njana/OneDrive/Desktop/web/webapp1/webapp/frontend/assets/" + header.Filename

		strId := r.Form.Get("mid")
		mob.ID, err = strconv.Atoi(strId)
		if err != nil {
			http.Error(w, "error parsing mob.id", http.StatusInternalServerError)
		}

		qry := `UPDATE products SET name = $1, specs = $2, price = $3, image_url = $4 WHERE id = $5`
		_, err = db.Exec(qry, mob.Name, mob.Specs, mob.Price, mob.Ipath, mob.ID)
		if err != nil {
			http.Error(w, "error updating mob "+err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(mob)

	}
}
