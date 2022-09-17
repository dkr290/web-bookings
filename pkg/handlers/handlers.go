package handlers

import (
	"net/http"

	"github.com/dkr290/web-bookings/pkg/config"
	"github.com/dkr290/web-bookings/pkg/models"
	"github.com/dkr290/web-bookings/pkg/render"
)

// Repo the repository used by handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// creates new repository
func NewRepo(a *config.AppConfig) *Repository {
	r := Repository{
		App: a,
	}
	return &r
}

// new handlers it sets repository for handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// this is the about page handler
func (r *Repository) About(w http.ResponseWriter, req *http.Request) {

	//perform some business login to the template
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again"

	remoteIP := r.App.Session.GetString(req.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	//send data to the template
	render.RenderTemplate(w, "about.page.go.html", &models.TemplateData{
		StringMap: stringMap,
	})

}

// this is home page handler
func (r *Repository) Home(w http.ResponseWriter, req *http.Request) {

	remoteIP := req.RemoteAddr
	r.App.Session.Put(req.Context(), "remote_ip", remoteIP)
	render.RenderTemplate(w, "index.page.go.html", &models.TemplateData{})
}
