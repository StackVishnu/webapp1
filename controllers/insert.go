package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"webapp/structures"
)

func Insert(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var mob structures.Mobile
		if err := parseForm(r, &mob); err != nil {
			errHandl(w, err, "Error parsing form data", http.StatusBadRequest)
			return
		}

		if err := saveFile(r, &mob); err != nil {
			errHandl(w, err, "Error saving file", http.StatusBadRequest)
			return
		}

		if err := insertIntoDB(db, mob); err != nil {
			errHandl(w, err, "Error inserting data", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Product inserted successfully")
	}
}

func parseForm(r *http.Request, mob *structures.Mobile) error {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		return err
	}
	mob.Name = r.FormValue("phone_name")
	mob.Specs = r.FormValue("specs")
	mob.Price, err = strconv.Atoi(r.Form.Get("price"))
	return err
}

func saveFile(r *http.Request, mob *structures.Mobile) error {
	file, header, err := r.FormFile("image")
	if err != nil {
		return err
	}
	defer file.Close()
	mob.Ipath = "assets/" + header.Filename // Set the image path in the mobile struct
	return nil
}

func insertIntoDB(db *sql.DB, mob structures.Mobile) error {
	qry := `INSERT INTO products (name, specs, price, image_url) VALUES ($1, $2, $3, $4)`
	_, err := db.Exec(qry, mob.Name, mob.Specs, mob.Price, mob.Ipath)
	return err
}
