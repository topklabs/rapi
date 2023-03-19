package websocket

import (
	"context"
	"github.com/gorilla/websocket"
	"time"
)

// Proxy is a bidirectional websocket proxy between a client and a server
type Proxy struct {
	client, server *websocket.Conn
}

// Run starts proxying connections between client and server.
func (p *Proxy) Run(ctx context.Context) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	p.run(ctx)
}

// Proxy messages from client to server.
func (p *Proxy) run(ctx context.Context) {
	errC := make(chan error, 1)

	for {
		select {
		case <-ctx.Done():
			p.shutdown(errC)
			return
		}
	}
}

func (p *Proxy) shutdown(errC chan error) {
	var err error

	now := time.Now()
	for _, conn := range []*websocket.Conn{p.client, p.server} {
		err = conn.SetReadDeadline(now)
		if err != nil {
			_ = p.client.Close()
		}
	}

	errC <- nil
}
