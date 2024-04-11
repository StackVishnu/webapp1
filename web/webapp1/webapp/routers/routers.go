package routers

import (
	"database/sql"
	"webapp/controllers"

	"github.com/gorilla/mux"
)

func Routes(db *sql.DB) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/insert", controllers.Insert(db))
	r.HandleFunc("/delete", controllers.Delete(db))
	r.HandleFunc("/viewAll", controllers.ViewAll(db))
	r.HandleFunc("/edit", controllers.Edit(db))
	r.HandleFunc("/search", controllers.Search(db))
	r.HandleFunc("/viewEach/{id}", controllers.ViewEach(db))
	r.HandleFunc("/login", controllers.Login(db))

	return r
}
