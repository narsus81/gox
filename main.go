package gox

import (
	"fmt"
	"net/http"
)

type Gox struct {
	config Config
}

type Config struct {
	debug   bool
	version string
}

func InitGox() *Gox {
	c := Config{debug: true, version: "v0.0.23"}
	http.HandleFunc("/gox", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Here it is, %q!", r.URL.Path)
	})
	return &Gox{config: c}
}

func (g *Gox) Version() string {
	return g.config.version
}
