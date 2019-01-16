package web

import (
	"html/template"
	"net/http"
)

type tplParams struct {
	Browser string
	URL string
}


func TemplatesStart() {
	http.HandleFunc("/text_templates/", textTemplateHandler)
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
