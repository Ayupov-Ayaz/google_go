package web

import (
	"fmt"
	"net/http"
)

type Handler struct {
	Name string // просто название
}

// Реализуем интерфейс net/http/server:
func (h *Handler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(rw, "Name:", h.Name, "Url:", req.URL.String())
}

func StartServerHttp() {
	// Создаем обработчик тест
	testHandler := &Handler{Name: "Test"}
	http.Handle("/test/", testHandler) // префикс url-а /test/

	// Создаем обработчик root
	rootHandler := &Handler{Name: "Root"}
	http.Handle("/", rootHandler) //будет срабатывать если нету нужного обработчика

	fmt.Println("starting server at :8000")
	http.ListenAndServe(":8000", nil)
}