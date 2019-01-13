package web

import (
	"fmt"
	"net/http"
)

func getHandler(w http.ResponseWriter, r *http.Request){
	// 1 способ получить get параметры:
	key := r.URL.Query().Get("key")
	if key != "" {
		fmt.Fprintln(w,"'key' is", key)
	}
	// 2 способ получить get параметры:
	value := r.FormValue("value")
	if value != "" {
		fmt.Fprintln(w, "'value' is", value)
	}
}

func StartWebGet() {
	mux.HandleFunc("/get/", getHandler)
	StartServeMux(8000)
}
