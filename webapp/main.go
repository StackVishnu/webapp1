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

	c := cors.New(
		cors.Options{
			AllowedOrigins: []string{"*"},
			AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		})

	handler := c.Handler(r)
	http.ListenAndServe(":9090", handler)

}
