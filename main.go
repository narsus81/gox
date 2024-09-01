package gox

import (
	"net/http"

	m "github.com/narsus81/gox/internal/model"
)

type Gox struct {
	version  string
	config   m.Config
	mux      *http.ServeMux
	patterns map[string]m.Route
	Chain    []m.Middleware
}

func Init() *Gox {
	c := m.Config{
		Debug:       true,
		HTMX:        `<script src="https://unpkg.com/htmx.org@2.0.2" integrity="sha384-Y7hw+L/jvKeWIRRkqWYfPcvVxHzVzn5REgzbawhxAuQGwX1XWe70vji+VSeHOThJ" crossorigin="anonymous"></script>`,
		SSE:         `<script src="https://unpkg.com/htmx-ext-sse@2.2.2/sse.js" crossorigin="anonymous"></script>`,
		Autogen:     `hx-ext="sse" sse-connect="/sse" sse-swap="First"`,
		Autogen2:    `hx-ext="sse" sse-connect="/sse2" sse-swap="Second"`,
		DefaultTmpl: "templates/default.tmpl",
	}
	g := Gox{version: "v0.1.3", config: c, mux: http.NewServeMux(), patterns: make(map[string]m.Route)}
	g.loadChain()

	g.HandleFunc("root", "/{$}", func(w http.ResponseWriter, r *http.Request) {})

	return &g
}
