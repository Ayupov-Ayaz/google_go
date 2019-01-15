package web

import (
	"fmt"
	"net/http"
)

/**
 Обслуживанеи статических данных без поднятия лишних серверов на чистом Go
 */
func StaticInServerStart() {
	http.HandleFunc("/main/", httpMainHandler)
	staticHandler := http.StripPrefix("/data/",
						http.FileServer(http.Dir("./static")),
					)
	http.Handle("/data/", staticHandler)
	fmt.Println("started server :8000")
	http.ListenAndServe(":8000", nil)
	}

func httpMainHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(`
		Hello World! <br>
	  	<img src="/data/img/gopher.png"/>
	`))
}