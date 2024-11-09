package core

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"runtime"
	"time"
)

type Route struct {
	Pattern string
	Handler interface{}
}

var m runtime.MemStats

func setupRoutes(s *Server, routes []Route) error {
	router := http.NewServeMux()

	wrapper := func(fn ServerHandler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			// TODO: middlewares

			fn(s, w, r)
			runtime.ReadMemStats(&m)

			log.Printf(
				"INFO: %s %v %s - RAM %.2f MB",
				r.Method,
				time.Since(start),
				r.URL.Path,
				float64(m.Alloc)/1024/1024,
			)
		})
	}

	for _, v := range routes {
		log.Printf("INFO: Registering route %q", v.Pattern)

		switch h := v.Handler.(type) {
		case func(*Server, http.ResponseWriter, *http.Request):
			router.Handle(v.Pattern, wrapper(h))
		case http.HandlerFunc:
			router.Handle(v.Pattern, h)
		case http.Handler:
			router.Handle(v.Pattern, h)
		default:
			return fmt.Errorf("invalid handler for pattern %q: %v", v.Pattern, reflect.TypeOf(h))
		}
	}

	s.Handler = router
	return nil
}
