package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/dkr290/web-bookings/internal/config"
	"github.com/dkr290/web-bookings/internal/models"
	"github.com/dkr290/web-bookings/internal/render"
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
	}, req)

}

// this is home page handler
func (r *Repository) Home(w http.ResponseWriter, req *http.Request) {

	remoteIP := req.RemoteAddr
	r.App.Session.Put(req.Context(), "remote_ip", remoteIP)
	render.RenderTemplate(w, "home.page.go.html", &models.TemplateData{}, req)
}

func (r *Repository) Generals(w http.ResponseWriter, req *http.Request) {

	remoteIP := req.RemoteAddr
	r.App.Session.Put(req.Context(), "remote_ip", remoteIP)
	render.RenderTemplate(w, "generals.page.go.html", &models.TemplateData{}, req)
}

func (r *Repository) Contact(w http.ResponseWriter, req *http.Request) {

	remoteIP := req.RemoteAddr
	r.App.Session.Put(req.Context(), "remote_ip", remoteIP)
	render.RenderTemplate(w, "contact.page.go.html", &models.TemplateData{}, req)
}
func (r *Repository) Majors(w http.ResponseWriter, req *http.Request) {

	remoteIP := req.RemoteAddr
	r.App.Session.Put(req.Context(), "remote_ip", remoteIP)
	render.RenderTemplate(w, "majors.page.go.html", &models.TemplateData{}, req)
}

// renders of search availability
func (r *Repository) Availability(w http.ResponseWriter, req *http.Request) {

	render.RenderTemplate(w, "search-availability.page.go.html", &models.TemplateData{}, req)
}
func (r *Repository) PostAvailability(w http.ResponseWriter, req *http.Request) {
	start := req.Form.Get("start")
	end := req.Form.Get("end")

	w.Write([]byte(fmt.Sprintf("Start date is %s and end date is %s", start, end)))
}

type jsonResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

// AvailabilityJSON Handles request for availability and sends JSOBN response
func (r *Repository) AvailabilityJSON(w http.ResponseWriter, req *http.Request) {
	resp := jsonResponse{
		OK:      true,
		Message: "Available",
	}

	out, err := json.MarshalIndent(resp, "", "     ")
	if err != nil {
		log.Println(err)
	}

	log.Println(string(out))
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)

}
func (r *Repository) Reservation(w http.ResponseWriter, req *http.Request) {

	remoteIP := req.RemoteAddr
	r.App.Session.Put(req.Context(), "remote_ip", remoteIP)
	render.RenderTemplate(w, "make-reservation.page.go.html", &models.TemplateData{}, req)
}
