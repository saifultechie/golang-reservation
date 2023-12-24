package main

import (
	"fmt"
	"github/saifultechie/booking/pkg/config"
	"testing"

	"github.com/go-chi/chi/v5"
)

func TestRoutes(t *testing.T) {
	var app config.AppConfig

	mux := Routes(&app)
	switch v := mux.(type) {
	case *chi.Mux:
		// do nothing

	default:
		t.Errorf(fmt.Sprintf("type is not http.handler but type is %T", v))

	}
}
