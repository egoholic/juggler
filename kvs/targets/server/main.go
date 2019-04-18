package main

import (
	"net"
	"net/http"

	"github.com/egoholic/juggler/kvs/config"
	"github.com/egoholic/juggler/kvs/node"
)

func main() {
	app := newApp()
	c, err := net.Dial("unix", config.UnixSocketPath())
	if err != nil {
		panic(err)
	}
	defer c.Close()

}

func newApp() http.Handler {
	return &App{node.New(config.ID())}
}

type App struct {
	node *node.Node
}
