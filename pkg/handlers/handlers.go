package handlers

import (
	"encoding/json"
	"fmt"
	"github/saifultechie/booking/pkg/config"
	"github/saifultechie/booking/pkg/forms"
	"github/saifultechie/booking/pkg/renders"
	models "github/saifultechie/booking/pkg/templateData"
	"net/http"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(a *Repository) {
	Repo = a
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	renders.RenderTemplate(w, r, "home.html", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// some business logic to generate data
	stringMap := make(map[string]string)
	stringMap["test"] = "hello world"

	remoteIp := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIp
	renders.RenderTemplate(w, r, "about.html", &models.TemplateData{StringMap: stringMap})
}

func (m *Repository) General(w http.ResponseWriter, r *http.Request) {
	// some business logic to generate data
	stringMap := make(map[string]string)
	renders.RenderTemplate(w, r, "general.html", &models.TemplateData{StringMap: stringMap})
}

func (m *Repository) Major(w http.ResponseWriter, r *http.Request) {
	// some business logic to generate data
	stringMap := make(map[string]string)
	renders.RenderTemplate(w, r, "major.html", &models.TemplateData{StringMap: stringMap})
}

func (m *Repository) Search(w http.ResponseWriter, r *http.Request) {
	// some business logic to generate data
	stringMap := make(map[string]string)
	renders.RenderTemplate(w, r, "search.html", &models.TemplateData{StringMap: stringMap})
}

func (m *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {
	// some business logic to generate data
	start := r.Form.Get("start_date")
	end := r.Form.Get("end_date")

	w.Write([]byte(fmt.Sprintf("start date is %s and end date is %s", start, end)))
}

func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	// some business logic to generate data
	var emptyReservation models.Reservation
	data := make(map[string]interface{})
	data["reservation"] = emptyReservation

	renders.RenderTemplate(w, r, "reservation.html", &models.TemplateData{
		Data: data,
		Form: forms.New(nil),
	})
}

// data parse from req.body
// at first we have to parse the request such as r.ParseForm()
//

func (m *Repository) ReservationPost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println(err)
		return
	}

	reservation := models.Reservation{
		FirstName: r.Form.Get("first_name"),
		LastName:  r.Form.Get("last_name"),
		Email:     r.Form.Get("email"),
		Phone:     r.Form.Get("phone"),
	}

	form := forms.New(r.PostForm)
	fmt.Println(r.PostForm)
	form.Required("first_name", "last_name", "email")
	form.MinLength("first_name", 3)
	form.IsEmail("email")

	if !form.Valid() {
		data := make(map[string]interface{})
		data["reservation"] = reservation
		renders.RenderTemplate(w, r, "reservation.html", &models.TemplateData{
			Data: data,
			Form: form,
		})
		return
	}
	// put data into session
	m.App.Session.Put(r.Context(), "reservation", reservation)
	http.Redirect(w, r, "/reservation-summary", http.StatusSeeOther)
}

func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	// some business logic to generate data
	stringMap := make(map[string]string)
	renders.RenderTemplate(w, r, "contact.html", &models.TemplateData{StringMap: stringMap})
}

type responseJson struct {
	Ok      bool   `json:"ok"`
	Message string `json:"message"`
}

func (m *Repository) AvailabilityJson(w http.ResponseWriter, r *http.Request) {
	resp := responseJson{
		Ok:      true,
		Message: "availability",
	}

	out, err := json.MarshalIndent(resp, "", "   ")
	if err != nil {
		fmt.Println("error when casting to json data")
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

func (m *Repository) ReservationSummary(w http.ResponseWriter, r *http.Request) {
	// some business logic to generate data
	reservation, ok := m.App.Session.Get(r.Context(), "reservation").(models.Reservation)
	if !ok {
		fmt.Println("can not get data from session")
		m.App.Session.Put(r.Context(), "error", "can't get reservation from storage")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	}
	// remove session after get data from session
	m.App.Session.Remove(r.Context(), "reservation")
	data := make(map[string]interface{})
	data["reservation"] = reservation
	renders.RenderTemplate(w, r, "reservation-summary.html", &models.TemplateData{
		Data: data,
	})
}
