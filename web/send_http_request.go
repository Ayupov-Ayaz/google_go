package web

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func SendHttpRequestToServer() {
	http.HandleFunc("/sendHttp", sendHttpRequestHandler)
	http.HandleFunc("/articles/", getHttpRequestHandler)
	fmt.Println("server started :8000")
	http.ListenAndServe(":8000", nil)
}

func sendHttpRequestHandler(w http.ResponseWriter, r *http.Request) {
	url := "http://localhost:8000/post/?id=123"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Ошибка при попытке получить Url:", err.Error())
	}
	defer resp.Body.Close() // Никогда не забывать!
	respBody, err := ioutil.ReadAll(resp.Body) // открыли файл
	fmt.Println(string(respBody))
}

func getHttpRequestHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "getHandler: incoming request %#v\n", r)
	fmt.Fprintln(w, "getHandler: r.Url %#v\n", r.URL)
}


