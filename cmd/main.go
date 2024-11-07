package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"gustavonovaes.dev/go-htmx/internal"
	"gustavonovaes.dev/go-htmx/internal/core"
	"gustavonovaes.dev/go-htmx/internal/count"
	"gustavonovaes.dev/go-htmx/internal/home"
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
	listenWithGracefulShutdown(addr, server)
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

	indexController := core.NewIndexController()
	homeController := home.NewHomeController()
	countController := count.NewCountController()

	routes := []core.Route{
		// Home
		{Pattern: "GET /home", Handler: homeController.RenderHome},

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
		{Pattern: "GET /", Handler: indexController.RenderIndex},
	}

	return core.NewServer(tr, routes)
}

func listenWithGracefulShutdown(addr string, server *core.Server) {
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		log.Printf("INFO: Listening in %s...", addr)
		err := http.ListenAndServe(addr, server)
		if err != nil {
			log.Fatalf("ERROR: Fail to start on addr: %q", addr)
		}
	}()

	<-stopChan
	log.Println("INFO: Shutting down server...")

	// Add your cleanup code here

	log.Println("INFO: Server gracefully stopped")
}
