package web

import (
	"fmt"
	"net/http"
	"time"
)



func CookiesStart() {
	mux.HandleFunc("/login", showLoginFormHandler)
	mux.HandleFunc("/login/in", loginHandler)
	mux.HandleFunc("/logout", logoutHandler)
	mux.HandleFunc("/main", mainHandler)
	StartServeMux(8000)
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	// проверка cookies :TODO не работает исправить
	login, err := r.Cookie("user_login")
	fmt.Println(r.Cookies(), err)
	if login != nil {
		fmt.Fprintln(w,"<a href=\"/logout\">Exit</a> <p>Hello,", login.Value)
		return
	}
	fmt.Fprintln(w,"<a href=\"/login\">Авторизоваться</a>")
}

func showLoginFormHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Write(loginFormTpl)
	}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	rightLogin := "tommy"
	rightPassword := "qwerty"
	err := r.ParseForm()
	if err !=  nil {
		fmt.Fprintln(w,"Не удалось распарсить форму")
	}
	login := r.Form.Get("login")
	password := r.Form.Get("password")

	if login == rightLogin && password == rightPassword {
		cookie := http.Cookie{
			Name: "user_login",
			Value: login,
			Expires: time.Now().Add(10 * time.Hour),
		}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/main", http.StatusFound)
	} else {
		http.Redirect(w, r, "/login", http.StatusFound)
	}
}



func logoutHandler(w http.ResponseWriter, r *http.Request) {
	user_login, _ := r.Cookie("user_login")
	if user_login == nil {
		http.Redirect(w, r, "/main", http.StatusOK)
		return
	}

	user_login.Expires = time.Now().AddDate(0, 0, -1)
	http.SetCookie(w, user_login)
}
