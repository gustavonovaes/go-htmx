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
	"gustavonovaes.dev/go-htmx/internal/core/middleware"
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

	homeController := home.NewHomeController()
	countController := count.NewCountController()

	withLog := middleware.CreateStack(
		middleware.Logger,
	)

	withUser := middleware.CreateStack(
		middleware.Logger,
		middleware.EnsureLoggedIn,
	)

	withAdmin := middleware.CreateStack(
		middleware.Logger,
		middleware.EnsureLoggedIn,
		middleware.EnsureAdmin,
	)

	return core.NewServer(tr, func(s *core.Server) http.Handler {
		wrapRoute := func(sh core.ServerHandler, middleware middleware.Middleware) http.Handler {
			handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				sh(s, w, r)
			})

			if middleware != nil {
				return middleware(handler)
			}

			return handler
		}

		routerHome := http.NewServeMux()
		routerHome.Handle("GET /", wrapRoute(homeController.RenderHome, withUser))

		routerCount := http.NewServeMux()
		routerCount.Handle("GET /count", wrapRoute(countController.RenderCount, withAdmin))
		routerCount.Handle("POST /count", wrapRoute(countController.IncCount, withAdmin))

		router := http.NewServeMux()
		router.Handle("/home", routerHome)
		router.Handle("/count", routerCount)

		router.Handle("GET /favicon.ico", http.NotFoundHandler())
		router.Handle(
			"GET /static/",
			http.StripPrefix("/static/", http.FileServer(http.Dir("static"))),
		)
		router.Handle("/", wrapRoute(core.NewIndexController().RenderIndex, withLog))

		return router
	})
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
