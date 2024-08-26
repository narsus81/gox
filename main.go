package gox

import (
	"fmt"
	"html/template"
	"net/http"
)

type Config struct {
	HTMX        template.HTML
	SSE         template.HTMLAttr
	defaultTmpl string
	debug       bool
}

type Gox struct {
	version  string
	config   Config
	mux      *http.ServeMux
	patterns map[string]Route
}

type Route struct {
	Name    string
	Handler func(http.ResponseWriter, *http.Request)
}

func (g *Gox) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// First we use a FakeWriter to get the pattern compiled by the mux
	// inefficient but it's the only way to get the chosen pattern
	fake := NewFakeResponseWriter()
	g.mux.ServeHTTP(fake, r)

	fmt.Println("Pattern: ", r.Pattern)
	fmt.Println("Method: ", r.Method)

	// Then we use the pattern to get the route
	if route, ok := g.patterns[r.Pattern]; !ok {
		http.NotFound(w, r)
		return
	} else {
		fmt.Println("Route found: ", route.Name, " with Pattern: ", r.Pattern)

		tmpl, err := template.New("default.tmpl").ParseGlob("templates/*.tmpl")
		if err != nil {
			panic(err)
		}

		err = tmpl.Execute(w, g.config)
		if err != nil {
			panic(err)
		}
	}

	// Real writer to serve the request
	g.mux.ServeHTTP(w, r)
}

func (g *Gox) HandleFunc(name string, pattern string, handler func(http.ResponseWriter, *http.Request)) {
	r := Route{Name: name, Handler: handler}
	g.patterns[pattern] = r
	g.mux.HandleFunc(pattern, handler)
}

func Init() *Gox {
	c := Config{
		debug:       true,
		HTMX:        `<script src="https://unpkg.com/htmx.org@1.9.12/dist/ext/sse.js"></script>`,
		SSE:         `hx-ext="sse" sse-connect="/sse" sse-swap="message"`,
		defaultTmpl: "default.tmpl",
	}
	g := Gox{version: "v0.0.25", config: c, mux: http.NewServeMux(), patterns: make(map[string]Route)}

	g.HandleFunc("root", "/", func(w http.ResponseWriter, r *http.Request) {})

	return &g
}

func (g *Gox) PrintVersion() string {
	return g.version
}
