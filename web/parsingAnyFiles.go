package web

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Params struct {
	Id int
	User string
}

func ParsingAnyFiles() {
	http.HandleFunc("/", uploadRawBody)
	http.ListenAndServe(":8000", nil)
}

func uploadRawBody(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	p := &Params{}

	err = json.Unmarshal(body, p)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	// извлекаем данные
	fmt.Fprintln( w,"Content-type-", r.Header.Get("Content-type"))
	fmt.Fprintln(w, "params: ", p)
}