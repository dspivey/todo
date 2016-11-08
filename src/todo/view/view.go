package view

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"todo/log"
)

var funcMap map[string]interface{} = template.FuncMap{
	"ToUpper": strings.ToUpper,
	"ToLower": strings.ToLower,
}

func getViewsDir() (path string) {
	cwd, _ := os.Getwd()
	return filepath.Join(cwd, "src", "todo", "view")
}

func ParseTemplates(templateNames ...string) (t *template.Template) {
	var files []string

	t = template.New("view").Funcs(funcMap)

	for _, file := range templateNames {
		fn := filepath.Join(getViewsDir(), fmt.Sprintf("%s.html", file))
		files = append(files, fn)
		log.Info(fn)
	}

	t = template.Must(t.ParseFiles(files...))

	return t
}

func RenderHTML(writer http.ResponseWriter, model interface{}, templateNames ...string) {
	var files []string

	for _, file := range templateNames {
		fn := filepath.Join(getViewsDir(), fmt.Sprintf("%s.html", file))
		files = append(files, fn)
		log.Info(fn)
	}

	templates := template.New("view").Funcs(funcMap)
	templates = template.Must(templates.ParseFiles(files...))

	err := templates.ExecuteTemplate(writer, "view", model)
	if err != nil {
		log.Danger("Error processing template: ", err)
	}
}
