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
	// r.HandleFunc("/insert", func(w http.ResponseWriter, r *http.Request) {
	// 	controllers.Insert(db, w, r)
	// 	fmt.Println("insert working")
	// })
	// r.HandleFunc("/delete", func(w http.ResponseWriter, r *http.Request) {
	// 	controllers.Delete(db, w, r)
	// 	fmt.Println("delete working")
	// })
	return r
}
