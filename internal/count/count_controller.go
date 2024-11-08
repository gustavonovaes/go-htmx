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

func (c *CountController) RenderCount(s *core.Server, w http.ResponseWriter, _ *http.Request) {
	s.Render(w, "count.html", map[string]interface{}{
		"Title": "HTMX & Go - Count",
		"Count": c.count.Count,
	})
}

func (c *CountController) GetCount(s *core.Server, w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Has("sleep") {
		ms, _ := strconv.Atoi(r.URL.Query().Get("sleep"))
		time.Sleep(time.Duration(ms) * time.Millisecond)
	}

	s.SendJSON(w, c.count)
}

func (c *CountController) IncCount(s *core.Server, w http.ResponseWriter, _ *http.Request) {
	c.count.Count++

	if (c.count.Count % 10) == 0 {
		s.Send(w, http.StatusInternalServerError, "Ops! Something went wrong")
		return
	}

	if (c.count.Count % 5) == 0 {
		s.Send(w, http.StatusTeapot, "I'm a teapot")
		return
	}

	s.SendJSON(w, c.count)
}

func NewCountController() *CountController {
	return &CountController{
		count: Count{
			Count: 0,
		},
	}
}
