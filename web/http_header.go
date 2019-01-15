package web

import (
	"fmt"
	"net/http"
)

func HttpHeaderStart() {
	mux.HandleFunc("/main/", MainHandler)
	StartServeMux(8000)
}

func MainHandler(w http.ResponseWriter, r *http.Request) {
	// Запись Header
	w.Header().Set("token", "ghg2323fjfj2323fkr42titi2vmvn234")

	// получение headers
	fmt.Fprintln(w, "token : ", w.Header().Get("token"))
	fmt.Fprintln(w, "browser", r.UserAgent())
	fmt.Fprintln(w, "your accept", r.Header.Get("Accept"))

}
