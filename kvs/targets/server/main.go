package main

import (
	"net/http"

	"github.com/egoholic/juggler/kvs/config"
)

func main() {
	app := NewApp()
	http.ListenAndServe(config.ListenTo(), app)
}

func NewApp() http.Handler {
  
}
