package handlers

import (
	"net/http"

	"github.com/devj1003/mygoapp/pkg/config"
	"github.com/devj1003/mygoapp/pkg/models"
	"github.com/devj1003/mygoapp/pkg/render"
)

// Repo the repository used for the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// HomePage func for home page handler
func (m *Repository) HomePage(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.html", &models.TemplateData{})
}

// AboutPage func for about page handler
func (m *Repository) AboutPage(w http.ResponseWriter, r *http.Request) {

	// perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again!"

	// send the data to the template
	render.RenderTemplate(w, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})

}
