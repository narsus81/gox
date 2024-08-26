package gox

import (
	"fmt"
	"html/template"
	"net/http"
)

type Middleware func(http.Handler) http.Handler

var chain []Middleware

func (g *Gox) loadChain() {
	chain = append(chain, g.basicMiddleware)
	chain = append(chain, g.rootMiddleware)
}

func chainingMiddleware(h http.Handler, m ...Middleware) http.Handler {
	if len(m) < 1 {
		return h
	}

	wrappedHandler := h
	for i := len(m) - 1; i >= 0; i-- {
		wrappedHandler = m[i](wrappedHandler)
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
			templ := template.Must(template.ParseFiles(g.config.defaultTmpl))
			templ.Execute(w, g.config)
		}
		next.ServeHTTP(w, r)
	})
}
