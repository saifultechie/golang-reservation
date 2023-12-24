package models

import "github/saifultechie/booking/pkg/forms"

type TemplateData struct {
	StringMap map[string]string
	IntMap    map[int]int
	Data      map[string]interface{}
	CSRFToken string
	Success   string
	Error     string
	Warning   string
	Flash     string
	Form      *forms.Form
}
