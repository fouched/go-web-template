package handlers

import (
	"github.com/fouched/go-web-template/internal/models"
	"github.com/fouched/go-web-template/internal/render"
	"net/http"
)

// Home is the home page handler
func (m *HandlerConfig) Home(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "/home.gohtml", &models.TemplateData{})
}
