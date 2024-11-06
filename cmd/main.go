package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"gustavonovaes.dev/go-htmx/internal"
	"gustavonovaes.dev/go-htmx/internal/core"
	"gustavonovaes.dev/go-htmx/internal/count"
)

const (
	PORT_DEFAULT = "8080"
)

func main() {
	server := createServer()

	port := os.Getenv("PORT")
	if port == "" {
		port = PORT_DEFAULT
	}

	addr := fmt.Sprintf(":%s", port)
	log.Printf("INFO: Listening in %s...", addr)
	err := http.ListenAndServe(addr, server)
	if err != nil {
		log.Fatalf("Fail to start on addr: %q", addr)
	}
}

func createServer() *core.Server {
	templates := []string{
		"*/views/*.html",
	}

	tr, err := internal.NewTemplateRenderer(templates...)
	if err != nil {
		log.Fatalf("ERROR: Fail to create template renderer: %v", err)
	}

	for _, name := range tr.GetTemplates() {
		log.Printf("INFO: Template found: %q", name)
	}

	countController := count.NewCountController()

	routes := []core.Route{
		// Count
		{Pattern: "GET /count", Handler: countController.RenderCount},
		{Pattern: "GET /api/count", Handler: countController.GetCount},
		{Pattern: "POST /api/count", Handler: countController.IncCount},

		// Static files
		{Pattern: "GET /favicon.ico", Handler: http.NotFoundHandler()},
		{
			Pattern: "GET /static/",
			Handler: http.StripPrefix("/static/", http.FileServer(http.Dir("static"))),
		},

		// Need to be the last route
		{Pattern: "GET /", Handler: indexHandler},
	}

	return core.NewServer(tr, routes)
}

func indexHandler(s *core.Server, w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		s.Send(w, http.StatusNotFound, "Page not found")
		return
	}

	s.Render(w, "index.html", nil)
}
