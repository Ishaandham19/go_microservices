package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Simple handler (satisfies http.Handler interface)
type Hello struct {
	l *log.Logger
}

// NewHello creates a new hello handler with the given logger
func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

// ServeHTTP implements the go http.Handler interface
func (h *Hello) ServeHTTP(rw http.ResponseWriter, r http.Request) {
	h.l.Println("Hello World")
	// read the body of req
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "Unable to read request body", http.StatusBadRequest)
		return
	}

	// write the response
	fmt.Fprintf(rw, "Hello %s", b)
}
