package renders

import (
	"bytes"
	"fmt"
	"github/saifultechie/booking/pkg/config"
	models "github/saifultechie/booking/pkg/templateData"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/justinas/nosurf"
)

var functions = template.FuncMap{}

var app *config.AppConfig

// set the new templates for the config
func NewTemplate(a *config.AppConfig) {
	app = a
}

func addTemplateData(td *models.TemplateData, r *http.Request) *models.TemplateData {
	td.Flash = app.Session.PopString(r.Context(), "flash")
	td.Error = app.Session.PopString(r.Context(), "error")
	td.Warning = app.Session.PopString(r.Context(), "warning")
	td.CSRFToken = nosurf.Token(r)
	return td
}

func RenderTemplate(w http.ResponseWriter, r *http.Request, tmpPath string, td *models.TemplateData) {
	var tc map[string]*template.Template
	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}
	t, ok := tc[tmpPath]
	if !ok {
		log.Fatal("can not get template from config")
	}

	buf := new(bytes.Buffer)
	td = addTemplateData(td, r)
	_ = t.Execute(buf, td)
	_, err := buf.WriteTo(w)

	if err != nil {
		fmt.Println("Error writing template to browser ", err)
	}
	// parsedTemplate, _ := template.ParseFiles("../../template/" + tmpPath)
	// err = parsedTemplate.Execute(w, nil)

	// if err != nil {
	// 	fmt.Println("parsing error ", err)
	// 	return
	// }
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("../../template/*.html")

	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("../../template/base.html")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("../../template/base.html")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}

	return myCache, nil
}
