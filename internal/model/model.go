package model

import (
	"html/template"
	"net/http"
)

type Config struct {
	HTMX        template.HTML
	SSE         template.HTML
	Autogen     template.HTMLAttr
	Autogen2    template.HTMLAttr
	DefaultTmpl string
	Debug       bool
}

type Route struct {
	Name    string
	Handler func(http.ResponseWriter, *http.Request)
}

type Middleware func(http.Handler) http.Handler
