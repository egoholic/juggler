package main

import (
	"io"
	"net/http"
)

var pipeline = pipeline.New()

func main() {
	http.ListenAndServe(":3333", &handler{})
}

type handler struct{}

func (handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var err error
	body, err := r.GetBody()
	if err != nil {
		// do nothing
	}

	logMsg := make([]byte, 1024)
	for err != io.EOF {
		_, err = body.Read(logMsg)
	}

	if err != nil && err != io.EOF {
		// do nothing
	}

	pipeline.Handle(logMsg)
}
