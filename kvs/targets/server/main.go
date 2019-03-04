package main

import (
	"net/http"

	"github.com/egoholic/juggler/kvs/config"
	"github.com/egoholic/juggler/kvs/node"
)

func main() {
	app := newApp()
	http.ListenAndServe(config.ListenTo(), app)
}

func newApp() http.Handler {
	return &App{node.New(config.ID())}
}

type App struct {
	node *node.Node
}

func (*App) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}
