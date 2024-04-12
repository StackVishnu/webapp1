package main

import (
	"net/http"
	"webapp/dbconn"
	"webapp/routers"

	_ "github.com/lib/pq"
	"github.com/rs/cors"
)

func main() {
	db := dbconn.Dbconnect()
	r := routers.Routes(db)
	// var store = sessions.NewCookieStore([]byte("your-secret-key"))
	c := cors.New(
		cors.Options{
			AllowedOrigins: []string{"*"},
			AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders: []string{"Content-Type", "Authorization"},
		})

	handler := c.Handler(r)
	http.ListenAndServe(":9090", handler)

}

// func Middleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

//         if r.URL.Path != "/login" {

//         }

//         next.ServeHTTP(w, r)
// 	})
// }
