package home

import (
	"net/http"

	"gustavonovaes.dev/go-htmx/internal/core"
)

type HomeController struct {
}

func (c *HomeController) RenderHome(s *core.Server, w http.ResponseWriter, _ *http.Request) {
	s.Render(w, "home.html", nil)
}

func NewHomeController() *HomeController {
	return &HomeController{}
}
