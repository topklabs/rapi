package websocket

import "github.com/gorilla/websocket"

// Conn wraps a pointer to a websocket connection.
type Conn struct {
	ws *websocket.Conn
}
