package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/dkr290/go-web/01-02-http-templates/pkg/config"
	"github.com/dkr290/go-web/01-02-http-templates/pkg/custerror"
	"github.com/dkr290/go-web/01-02-http-templates/pkg/models"
)

var app *config.AppConfig

//Newtemplates sets the confit to template package

func Newtemplates(a *config.AppConfig) {
	app = a

}

func AddDEfaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {

	//get the template cache from app config
	var tc map[string]*template.Template
	var err error
	if app.UseCache {

		tc = app.TemplateCache
	} else {
		tc, err = CreateTemplateCache()
		custerror.FatalErr(err)
	}

	//create template cache

	//commented out because it gets form app.config
	//tc, err := CreateTemplateCache()
	//custerror.FatalErr(err)

	//get reused template from cache

	// t, ok := tc[tmpl]
	// if !ok {
	// 	log.Fatal(err)
	// }

	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get template from template cache")
	}

	buf := new(bytes.Buffer)

	td = AddDEfaultData(td)
	err = t.Execute(buf, td)
	if err != nil {
		log.Fatal(err)
	}

	//render the template

	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {

	myCache := map[string]*template.Template{}

	//get all of the files page.gohtml

	pages, err := filepath.Glob("./templates/*.page.gohtml")
	custerror.FatalErr(err)

	for _, page := range pages {
		//	first we need to get the actual name of the page using filepath.Base
		name := filepath.Base(page)
		// next, we need to actually create the template set parsing the file as well
		ts, err := template.New(name).ParseFiles(page)
		custerror.FatalErr(err)

		//	Next, the template set needs to know of any layouts we are using so it can parse correctly

		matches, err := filepath.Glob("./templates/*.layout.gohtml")
		custerror.FatalErr(err)
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.gohtml")
			custerror.FatalErr(err)
		}

		myCache[name] = ts
	}

	return myCache, nil

}
