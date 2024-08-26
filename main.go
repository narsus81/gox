package gox

import (
	"fmt"
	"net/http"
)

type Gox struct {
	config Config
	mux   *http.ServeMux

}

type Config struct {
	debug   bool
	version string
}

func InitGox() *Gox {
	c := Config{debug: true, version: "v0.0.25"}
	g := Gox{config: c, mux: http.NewServeMux()}
	g.mux.handleFunc("/gox", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<h1>Hello from GOX, the HTMX framework</h1>\n you've requested: %s\n", r.URL.Path)
	})
	return &g
}

func (g *Gox) Version() string {
	return g.config.version
}
