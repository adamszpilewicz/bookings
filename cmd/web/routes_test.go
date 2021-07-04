package main

import (
	"github.com/adamszpilewicz/bookings/internal/config"
	"github.com/go-chi/chi"
	"testing"
)

func TestRoutes(t *testing.T) {
	var app config.AppConfig

	mux := routes(&app)

	switch mux.(type) {
	case *chi.Mux:
	default:
		t.Error("not type of *chi.Mux")
	}

}
