package count

import (
	"net/http"
	"strconv"
	"time"

	"gustavonovaes.dev/go-htmx/internal/core"
)

type CountController struct {
	count Count `json:"count"`
}

func (c *CountController) RenderCount(s *core.Server, w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Accept") == "application/json" {
		if r.URL.Query().Has("sleep") {
			ms, _ := strconv.Atoi(r.URL.Query().Get("sleep"))
			time.Sleep(time.Duration(ms) * time.Millisecond)
		}

		s.SendJSON(w, c.count, http.StatusOK)
		return
	}

	s.Render(w, "count.html", map[string]interface{}{
		"Count": c.count.Count,
	})
}

func (c *CountController) GetCount(s *core.Server, w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Has("sleep") {
		ms, _ := strconv.Atoi(r.URL.Query().Get("sleep"))
		time.Sleep(time.Duration(ms) * time.Millisecond)
	}

	s.SendJSON(w, c.count, http.StatusOK)
}

func (c *CountController) IncCount(s *core.Server, w http.ResponseWriter, _ *http.Request) {
	c.count.Count++

	if (c.count.Count % 10) == 0 {
		s.Send(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	if (c.count.Count % 5) == 0 {
		s.Send(w, http.StatusTeapot, http.StatusText(http.StatusTeapot))
		return
	}

	s.SendJSON(w, c.count, http.StatusOK)
}

func NewCountController() *CountController {
	return &CountController{
		count: Count{
			Count: 0,
		},
	}
}
