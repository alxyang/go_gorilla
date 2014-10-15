package hardware

import(
  "log"
)

type Hub struct {
  // Registered connections.
  connections map[*connection]bool

  // Inbound messages from the connections.
  broadcast chan []byte

  // Register requests from the connections.
  register chan *connection

  // Unregister requests from connections.
  unregister chan *connection
}

var H = Hub{
  broadcast:   make(chan []byte),
  register:    make(chan *connection),
  unregister:  make(chan *connection),
  connections: make(map[*connection]bool),
}

func (h *Hub) Run() {
  log.Println("Websocket hub is running.");
  for {
    select {
    case c := <-h.register:
      h.connections[c] = true
    case c := <-h.unregister:
      if _, ok := h.connections[c]; ok {
        delete(h.connections, c)
        close(c.send)
      }
    case m := <-h.broadcast:
      for c := range h.connections {
        select {
        case c.send <- m:
        default:
          delete(h.connections, c)
          close(c.send)
        }
      }
    }
  }
}
