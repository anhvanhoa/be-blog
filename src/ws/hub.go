package ws

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	clients map[string]*Client

	// Inbound messages from the clients.
	broadcast chan []byte

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client
}

func newHub() *Hub {
	return &Hub{
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[string]*Client),
	}
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			c, ok := h.clients[client.id]
			if !ok {
				h.clients[client.id] = client
			} else {
				c.tabCount++
			}
		case client := <-h.unregister:
			id := client.id
			if c, ok := h.clients[id]; ok {
				if c.tabCount > 1 {
					c.tabCount--
				} else {
					delete(h.clients, id)
					close(client.send)
				}
			}
		case message := <-h.broadcast:
			for id, client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, id)
				}
			}
		}

	}
}
