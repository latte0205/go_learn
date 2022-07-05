package handlers

import (
	"net/http"

	"github.com/tsawler/go-course/pkg/render"
)

// Home is the handler for the home page
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.html")
}

func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.html")
}
