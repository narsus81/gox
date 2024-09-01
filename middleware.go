package gox

import (
	"fmt"
	"html/template"
	"net/http"

	m "github.com/narsus81/gox/internal/model"
)

func (g *Gox) loadChain() {
	g.Chain = append(g.Chain, g.basicMiddleware)
	g.Chain = append(g.Chain, g.rootMiddleware)
}

func chainingMiddleware(h http.Handler, mdw ...m.Middleware) http.Handler {
	if len(mdw) < 1 {
		return h
	}

	wrappedHandler := h
	for i := len(mdw) - 1; i >= 0; i-- {
		wrappedHandler = mdw[i](wrappedHandler)
	}

	return wrappedHandler
}

func (g *Gox) basicMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		route := g.patterns[r.Pattern]
		fmt.Println("First middleware, the Pattern is: ", r.Pattern, " and the Route is: ", route.Name)
		next.ServeHTTP(w, r)
	})
}

func (g *Gox) rootMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		route := g.patterns[r.Pattern]
		if route.Name == "root" {
			fmt.Println("Root middleware, the Pattern is: ", r.Pattern, " and the Route is: ", route.Name)
			templ := template.Must(template.ParseFiles(g.config.DefaultTmpl))
			templ.Execute(w, g.config)
		}
		next.ServeHTTP(w, r)
	})
}
