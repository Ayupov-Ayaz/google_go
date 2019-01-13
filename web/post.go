package web

import (
	"fmt"
	"net/http"
)

func postHandler(w http.ResponseWriter, r *http.Request) {
	// 1 способ: Руками распарсить форму
	err := r.ParseForm()
	if err != nil {
		panic(err.Error())
	}
	login := r.Form["login"][0]
	password := r.Form["password"][0]
	fmt.Fprintln(w,"1 способ) Login:", login, "\t Password:", password)

	// 2 способ: Вытащить непосредственно из самой формы
	inputLogin := r.Form.Get("login")
	inputPassword := r.Form.Get("password")
	fmt.Fprintln(w,"2 способ) Login:", inputLogin, "\t Password:", inputPassword)
}

func showFormHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Write(loginFormTpl)
	}
}

func StartWebPost() {
	mux.HandleFunc("/login/in", postHandler)
	mux.HandleFunc("/login", showFormHandler)
	StartServeMux(8001)
}

var loginFormTpl = []byte(`
<html>
	<body>
		<form action="/login/in" method="post">
			Login:<input type="string" name="login" placeholder="| your login"><br>
			Password:<input type="password" name="password" placeholder="| your password"><br>
			<input type="submit" name="Отправить">
		</form>
	</body>	
</html>	
`)
