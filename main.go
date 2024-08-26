package gox

import (
	"html/template"
	"net/http"
)

type Config struct {
	HTMX        template.HTML
	SSE         template.HTML
	Autogen     template.HTMLAttr
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

func Init() *Gox {
	c := Config{
		debug:   true,
		HTMX:    `<script src="https://unpkg.com/htmx.org@2.0.2" integrity="sha384-Y7hw+L/jvKeWIRRkqWYfPcvVxHzVzn5REgzbawhxAuQGwX1XWe70vji+VSeHOThJ" crossorigin="anonymous"></script>`,
		SSE:     `<script src="https://unpkg.com/htmx-ext-sse@2.2.2/sse.js"></script>`,
		Autogen: `hx-ext="sse" sse-connect="/sse" sse-swap="message"`,
		//Autogen:     `hx-ext="sse" sse-connect="/sse"`,
		defaultTmpl: "templates/default.tmpl",
	}
	g := Gox{version: "v0.0.27", config: c, mux: http.NewServeMux(), patterns: make(map[string]Route)}
	g.loadChain()

	g.HandleFunc("root", "/{$}", func(w http.ResponseWriter, r *http.Request) {})

	return &g
}

func (g *Gox) HandleFunc(name string, pattern string, handler func(http.ResponseWriter, *http.Request)) {
	// Add the route to the patterns map
	r := Route{Name: name, Handler: handler}
	g.patterns[pattern] = r
	// Wrap the handler with the default middleware
	g.mux.Handle(pattern, chainingMiddleware(http.HandlerFunc(handler), chain...))
}

func (g *Gox) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	g.mux.ServeHTTP(w, r)
}
