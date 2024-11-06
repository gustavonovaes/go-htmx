package count

import (
	"net/http"

	"gustavonovaes.dev/go-htmx/internal/core"
)

type CountController struct {
	count Count `json:"count"`
}

func (c *CountController) RenderCount(s *core.Server, w http.ResponseWriter, _ *http.Request) {
	s.Render(w, "count.html", map[string]interface{}{
		"Title": "HTMX & Go - Count",
		"Count": c.count.Count,
	})
}

func (c *CountController) GetCount(s *core.Server, w http.ResponseWriter, _ *http.Request) {
	s.SendJSON(w, c.count)
}

func (c *CountController) IncCount(s *core.Server, w http.ResponseWriter, _ *http.Request) {
	c.count.Count++
	s.SendJSON(w, c.count)
}

func NewCountController() *CountController {
	return &CountController{
		count: Count{
			Count: 0,
		},
	}
}
