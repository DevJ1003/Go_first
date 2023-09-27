package handlers

import (
	"net/http"

	"github.com/devj1003/mygoapp/pkg/render"
)

// HomePage func for home page handler
func HomePage(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.html")
}

// AboutPage func for about page handler
func AboutPage(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.html")
}
