package core

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"gustavonovaes.dev/go-htmx/internal"
)

type ServerHandler func(s *Server, w http.ResponseWriter, r *http.Request)

type Server struct {
	http.Handler     // Embedding the http.Handler interface
	templateRenderer *internal.TemplateRenderer
}

func (s *Server) Render(w http.ResponseWriter, template string, data interface{}) {
	err := s.templateRenderer.Render(w, template, data)
	if err == nil {
		return
	}

	ErrorHandler(w, fmt.Errorf("ERROR: Fail to render template: %q, %v", template, err))
}

func (s *Server) Send(w http.ResponseWriter, statusCode int, content string) error {
	return writeResponse(w, statusCode, content, "text/html")
}

func (s *Server) SendJSON(w http.ResponseWriter, content interface{}) error {
	return writeResponse(w, http.StatusOK, content, "application/json")
}

func NewServer(t *internal.TemplateRenderer, routes []Route) *Server {
	var server = new(Server)

	server.templateRenderer = t

	err := setupRoutes(server, routes)
	if err != nil {
		log.Fatalf("ERROR: Fail to setup routes: %v", err)
	}

	return server
}

func ErrorHandler(w http.ResponseWriter, err error) {
	log.Printf("ERROR: %v", err)
	err = writeResponse(
		w,
		http.StatusInternalServerError,
		http.StatusText(http.StatusInternalServerError),
		"text/html",
	)

	if err != nil {
		log.Printf("ERROR: Sending response to client: %v", err)
	}
}

func writeResponse(
	w http.ResponseWriter,
	statusCode int,
	content interface{},
	contentType string,
) error {
	w.WriteHeader(statusCode)
	switch contentType {
	case "application/json":
		w.Header().Set("content-type", "application/json")
		return json.NewEncoder(w).Encode(content)
	case "application/html":
	case "application/plain":
		w.Header().Set("content-type", contentType)
		_, err := fmt.Fprintln(w, content)
		return err
	default:
		w.Header().Set("content-type", "text/plain")
		_, err := fmt.Fprintln(w, content)
		return err
	}

	return nil
}
