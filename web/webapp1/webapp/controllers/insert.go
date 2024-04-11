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

func Insert(db *sql.DB) http.HandlerFunc {
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

		qry := `INSERT INTO products ("name", "specs", "price", "image_url") VALUES ($1, $2, $3, $4)`
		_, err = db.Exec(qry, mob.Name, mob.Specs, mob.Price, mob.Ipath)
		if err != nil {
			http.Error(w, "error inserting data"+err.Error(), http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(mob)

	}
}
