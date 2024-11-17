package core

import "net/http"

type IndexController struct {
}

func (c *IndexController) RenderIndex(s *Server, w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		if r.Header.Get("Accept") == "application/json" {
			s.SendJSON(w, map[string]interface{}{
				"error": http.StatusText(http.StatusNotFound),
			}, http.StatusNotFound)
			return
		}

		s.Send(w, http.StatusNotFound, http.StatusText(http.StatusNotFound))
		return
	}

	s.Render(w, "index.html", map[string]interface{}{
		"Title":       "HTMX & Go",
		"Description": "Learn to count with HTMX and Go",
	})
}

func NewIndexController() *IndexController {
	return &IndexController{}
}
