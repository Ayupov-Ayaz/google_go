package web

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

/************************************* TEXT TEMPLATE ************************************************/

type tplParams struct {
	Browser string
	URL string
}

func TemplatesStart() {
	http.HandleFunc("/text_templates/", textTemplateHandler)
	http.HandleFunc("/html_templates/", httpTemplateHandler)
	http.ListenAndServe(":8000", nil)
}

var text = `
	Your browser: {{.Browser}}
	URL : {{.URL}}
`

// TextTemplate всё выводи как есть, так как ему передашь
func textTemplateHandler(w http.ResponseWriter, r *http.Request) {
	textTempl := template.New("text_template_example")
	textTempl, _ = textTempl.Parse(text)
	params := tplParams{
		Browser: r.UserAgent(),
		URL: r.URL.String(),
	}

	textTempl.Execute(w, params)
}

/************************************* HTML TEMPLATE ************************************************/
type User struct {
	ID int
	Name string
	Active bool
	Age uint
}

// в template можно вызывать пользовательские функции, обращаться к ним как к переменным без ()
func (u *User) IsActive() string{
	if u.Active {
		return "youser is active"
	}
	return ""
}

// Можно создать функции для шаблонизатора
func IsUserAdult(u *User) bool{
	if u.Age >= 18 {
		return true
	}
	return false
}

/**
    Добавляет необходимые экранирования:
	в URL он сделает URL escape
	в HTML сделает escape HTML сущностей и т.д.
  */
func httpTemplateHandler(w http.ResponseWriter, r *http.Request){
	// парсим наш файл
	//tmpl := template.Must(template.ParseFiles("templates/users.html")) //можно и так парсить, но mapFunc сюда не отправишь

	// Прокидываем нашу функцию в карту функций шаблона
	tmplFuncs := template.FuncMap{
		"isUserAdult" : IsUserAdult,
	}

	b, err := ioutil.ReadFile("templates/users.html")
	if err != nil {
		panic(err)
	}

	// Прокидываем funcMap, парсим наш шаблон
	tmpl, err := template.
		New("templates/users.html").Funcs(tmplFuncs).Parse(string(b))
	if err != nil {
		panic(err)
	}

	users := []User{
		{1, "Максим", true, 27},
		{1, "<i>Роман</i>", false, 17}, // Экранирует html тэги
		{1, "Антон", true, 14},
		{1, "Василий", true, 37},
	}

	// прокидываем нашу структуру
	err = tmpl.Execute(w, struct {
		Users []User // Создаем объъект slice к которому можно будет обратиться в шаблоне
	}{
		users, // что из себя представляет Users
	})

	if err != nil {
		fmt.Println(err.Error())
	}
}
