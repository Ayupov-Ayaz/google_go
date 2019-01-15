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
	user_login, err := r.Cookie("user_login")
	fmt.Println(user_login, err)
	logged := err == nil && user_login != nil
	if  logged {
		fmt.Fprintln(w,"<a href=\"/logout\">Exit</a> <p>Hello,", user_login.Value)
		return
	}
	fmt.Fprintln(w, "<a href=\"/login\">Авторизоваться</a>")
}

// показ формы или приветствие
func showLoginFormHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		user_login, err := r.Cookie("user_login")
		logged := err == nil && user_login != nil
		if !logged {
			w.Write(loginFormTpl)
		} else {
			fmt.Fprintln(w, "Привет, ", user_login.Value)
		}
	}
}

// авторизация
func loginHandler(w http.ResponseWriter, r *http.Request) {
	rightLogin := "ray"
	rightPassword := "43"
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
			Path: "/",
		}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/main", http.StatusFound)
		return
	}
	http.Redirect(w, r, "/login", http.StatusFound)
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	user_login, err := r.Cookie("user_login")
	if user_login == nil && err != nil {
		http.Redirect(w, r, "/main", http.StatusOK)
		return
	}

	user_login.Expires = time.Now().AddDate(-1, 0, 0)
	http.SetCookie(w, user_login)
	http.Redirect(w, r, "/main", http.StatusOK)
}
