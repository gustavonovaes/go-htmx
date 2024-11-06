package core

import "net/http"

type IndexController struct {
}

func (c *IndexController) RenderIndex(s *Server, w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		s.Send(w, http.StatusNotFound, "Page not found")
		return
	}

	s.Render(w, "index.html", nil)
}
