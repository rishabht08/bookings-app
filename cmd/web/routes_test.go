package main

import (
	"fmt"
	"testing"

	"github.com/go-chi/chi"
	"github.com/rishabht08/bookings-app/internal/config"
)

func TestRoutes(t *testing.T) {
	h := routes(&config.AppConfig{})

	switch v := h.(type) {
	case *chi.Mux:
		// do nothing
	default:
		t.Error(fmt.Sprintf("Type is not *chi.Mux , but is %T", v))
	}
}
