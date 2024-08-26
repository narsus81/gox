package gox

import (
	"io"
	"net/http"
)

// FakeResponseWriter is a fake implementation of http.ResponseWriter
type FakeResponseWriter struct {
}

func NewFakeResponseWriter() *FakeResponseWriter {
	return &FakeResponseWriter{}
}

func (r *FakeResponseWriter) Header() http.Header {
	return nil
}

func (r *FakeResponseWriter) Write(body []byte) (int, error) {
	return io.Discard.Write(body)
}

func (r *FakeResponseWriter) WriteHeader(status int) {
}
