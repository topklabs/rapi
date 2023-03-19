package websocket

import (
	"context"
	"github.com/gorilla/websocket"
	"net/http"
)

// Relay is sets up a websocket proxy between a client and a server.
type Relay struct {
	context  context.Context
	config   Config
	proxy    *Proxy
	upgrader websocket.Upgrader
}

// NewRelay stands up the configuration for the relay that will instantiate a proxy.
func NewRelay(ctx context.Context, cfg Config) *Relay {
	return &Relay{
		context:  ctx,
		config:   cfg,
		upgrader: websocket.Upgrader{},
	}
}

// Proxy stands up a proxy for the relay service.
func (rl *Relay) Proxy(client, server *websocket.Conn) {
	rl.proxy = &Proxy{client: client, server: server}
	rl.proxy.Run(rl.context)
}

// Do relays websocket requests from client to server and server to client.
func (rl *Relay) Do(w http.ResponseWriter, r *http.Request) {
	client, err := rl.upgrader.Upgrade(w, r, map[string][]string{})
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	server, _, err := websocket.DefaultDialer.Dial(rl.config.Host, map[string][]string{})
	if err != nil {
		_ = client.Close()
		return
	}

	rl.Proxy(client, server)
}
