package controllers

import "net/http"

func errHandl(w http.ResponseWriter, err error, str string, status int) {
	http.Error(w, str+err.Error(), status)
}
