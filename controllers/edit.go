package controllers

import (
	"database/sql"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"webapp/structures"
)

func Edit(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var mob structures.Mobile

		// Parse form data
		err := r.ParseMultipartForm(10 << 20)
		if err != nil {
			http.Error(w, "Error parsing form data", http.StatusBadRequest)
			return
		}

		// Parse mobile data
		mob.Name = r.Form.Get("phone_name")
		mob.Specs = r.Form.Get("specs")
		strPrice := r.Form.Get("price")
		mob.Price, err = strconv.Atoi(strPrice)
		if err != nil {
			http.Error(w, "Error parsing phone price", http.StatusBadRequest)
			return
		}

		// Retrieve image file
		file, header, err := r.FormFile("image")
		if err != nil || file == nil || header == nil {
			http.Error(w, "Error retrieving image file", http.StatusBadRequest)
			return
		}
		defer file.Close()

		// Create and save the image file
		out, err := os.Create("./frontend/assets/" + header.Filename)
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
		mob.Ipath = "assets/" + header.Filename

		// Parse mobile ID
		strID := r.Form.Get("mid")
		mob.ID, err = strconv.Atoi(strID)
		if err != nil {
			http.Error(w, "Error parsing mob.id", http.StatusInternalServerError)
			return
		}

		// Update the database record
		qry := `UPDATE products SET name = $1, specs = $2, price = $3, image_url = $4 WHERE id = $5`
		_, err = db.Exec(qry, mob.Name, mob.Specs, mob.Price, mob.Ipath, mob.ID)
		if err != nil {
			http.Error(w, "Error updating mob: "+err.Error(), http.StatusInternalServerError)
			return
		}

		// Return success response
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Mobile record updated successfully")
	}
}
