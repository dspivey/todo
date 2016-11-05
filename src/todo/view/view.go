package view

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
)

func getViewsDir() (path string) {
	cwd, _ := os.Getwd()
	return filepath.Join(cwd, "src", "todo", "view")
}

func ParseTemplates(templateNames ...string) (t *template.Template) {
	var files []string

	t = template.New("views")
	for _, file := range templateNames {
		fn := filepath.Join(getViewsDir(), fmt.Sprintf("%s.html", file))
		files = append(files, fn)
		fmt.Println(fn)
	}

	t = template.Must(t.ParseFiles(files...))

	return t
}

func RenderHTML(writer http.ResponseWriter, model interface{}, templateNames ...string) {
	var files []string

	for _, file := range templateNames {
		fn := filepath.Join(getViewsDir(), fmt.Sprintf("%s.html", file))
		files = append(files, fn)
		fmt.Println(fn)
	}

	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(writer, "views", model)
}
