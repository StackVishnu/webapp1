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
		err := r.ParseMultipartForm(10 << 20) // Set the maximum form size to 10 MB
		if err != nil {
			http.Error(w, "Error parsing form data", http.StatusBadRequest)
			return
		}

		// Get the form values
		strID := r.Form.Get("mid")
		fmt.Printf("strID")
		mob.ID, err = strconv.Atoi(strID)
		if err != nil {
			http.Error(w, "Error parsing mob.id", http.StatusInternalServerError)
			return
		}

		mob.Name = r.Form.Get("phone_name")
		mob.Specs = r.Form.Get("specs")
		strPrice := r.Form.Get("price")
		mob.Price, err = strconv.Atoi(strPrice)
		if err != nil {
			http.Error(w, "Error parsing phone price", http.StatusBadRequest)
			return
		}

		file, header, err := r.FormFile("image")
		if err != nil {
			http.Error(w, "Error retrieving image file", http.StatusBadRequest)
			return
		}
		defer file.Close()

		fmt.Println("Uploaded file:", header.Filename)
		mob.Ipath = "assets/" + header.Filename
		fmt.Printf("%v", mob)

		qry := `UPDATE products SET name = $1, specs = $2, price = $3, image_url = $4 WHERE id = $5`
		_, err = db.Exec(qry, mob.Name, mob.Specs, mob.Price, mob.Ipath, mob.ID)
		if err != nil {
			http.Error(w, "Error updating product: "+err.Error(), http.StatusInternalServerError)
			return
		}

		// Return success response
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Product updated successfully")
	}
}
