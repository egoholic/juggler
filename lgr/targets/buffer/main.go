package main

import (
	"net/http"

	"github.com/egoholic/juggler/lgr/config"
)

func main() {
	app := newApp()
	http.ListenAndServe(config.ListenTo(), app)
}

func newApp() http.Handler {
	return &App{}
}

type App struct{}

func (*App) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}
