package routers

import (
	"database/sql"
	"webapp/controllers"

	"github.com/gorilla/mux"
)

func Routes(db *sql.DB) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/insert", controllers.Insert(db)).Methods("POST")
	r.HandleFunc("/delete", controllers.Delete(db)).Methods("DELETE")
	r.HandleFunc("/viewAll", controllers.ViewAll(db)).Methods("GET")
	r.HandleFunc("/edit", controllers.Edit(db)).Methods("POST")
	r.HandleFunc("/search", controllers.Search(db)).Methods("POST")
	r.HandleFunc("/viewEach/{id}", controllers.ViewEach(db)).Methods("GET")
	r.HandleFunc("/login", controllers.Login(db)).Methods("POST")

	return r
}
