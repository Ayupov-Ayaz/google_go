package web

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "URL:", r.URL.String())
}
/**
	Создание сервера через мультиплексер предоставляет больше возможностей для настройки
 */
func StartServeMux(port int) {
	// Используется для поднятия нескольких разных урлов или несколько разных серверов
	mux := http.NewServeMux()
	//можем создать несколько обработчиков, которые обрабатывают одни и те же url-ы но по разному
	mux.HandleFunc("/", handler)
	sPort := ":" + strconv.Itoa(port)
	server := http.Server{
		Addr: sPort, // порт который сервер будет обрабатывать
		Handler: mux, // мультиплексор обработчик запросов
		ReadTimeout: 10 * time.Second, // время ожидания на чтение
		WriteTimeout: 10 * time.Second, // время ожидание на запись
	}
	fmt.Printf("Starting server at %s \n", sPort)
	server.ListenAndServe()
}