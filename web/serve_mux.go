package web

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "1 Handler, URL:", r.URL.String())
}

func StartServerMux(port int) {
	// Используется для поднятия нескольких разных урлов или несколько разных серверов
	mux := http.NewServeMux()
	//можем создать несколько обработчиков, которые обрабатывают одни и те же url-ы но по разному
	mux.HandleFunc("/", handler)
	sPort := ":" + strconv.Itoa(port)
	server := http.Server{
		Addr: sPort,
		Handler: mux,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	fmt.Printf("Starting server at %s", sPort)
	server.ListenAndServe()
}