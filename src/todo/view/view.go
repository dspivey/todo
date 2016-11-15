package view

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
	"todo/log"
)

var funcMap map[string]interface{} = template.FuncMap{
	"ToUpper":    strings.ToUpper,
	"ToLower":    strings.ToLower,
	"Now":        time.Now,
	"FormatDate": formatDate,
}

func formatDate(t time.Time, layout string) string {
	return t.Format(layout)
}

func getViewsDir() (path string) {
	cwd, _ := os.Getwd()
	return filepath.Join(cwd, "src", "todo", "view")
}

func ParseTemplates(templateName string, templateFiles ...string) (t *template.Template) {
	var files []string

	t = template.New(templateName).Funcs(funcMap)

	for _, file := range templateFiles {
		fn := filepath.Join(getViewsDir(), fmt.Sprintf("%s.html", file))
		files = append(files, fn)
		log.Info(fn)
	}

	t = template.Must(t.ParseFiles(files...))

	return t
}

func RenderHTML(writer http.ResponseWriter, model interface{}, templateName string, templateFiles ...string) {
	var files []string

	for _, file := range templateFiles {
		fn := filepath.Join(getViewsDir(), fmt.Sprintf("%s.html", file))
		files = append(files, fn)
		log.Info(fn)
	}

	templates := template.New(templateName).Funcs(funcMap)
	templates = template.Must(templates.ParseFiles(files...))

	err := templates.ExecuteTemplate(writer, templateName, model)
	if err != nil {
		log.Danger("Error processing template: ", err)
	}
}
