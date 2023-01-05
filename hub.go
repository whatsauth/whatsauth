package whatsauth

type message struct {
	data []byte
	room string
}

type subscription struct {
	conn *connection
	room string
}

// hub maintains the set of active connections and broadcasts messages to the
// connections.
type hub struct {
	// Registered connections.
	rooms map[string]map[*connection]bool

	// Inbound messages from the connections.
	broadcast chan message

	// Register requests from the connections.
	register chan subscription

	// Unregister requests from connections.
	unregister chan subscription
}

var Hub = hub{
	broadcast:  make(chan message),
	register:   make(chan subscription),
	unregister: make(chan subscription),
	rooms:      make(map[string]map[*connection]bool),
}

func HubRun(Hub *hub) {
	for {
		select {
		case s := <-Hub.register:
			connections := Hub.rooms[s.room]
			if connections == nil {
				connections = make(map[*connection]bool)
				Hub.rooms[s.room] = connections
			}
			Hub.rooms[s.room][s.conn] = true
		case s := <-Hub.unregister:
			connections := Hub.rooms[s.room]
			if connections != nil {
				if _, ok := connections[s.conn]; ok {
					delete(connections, s.conn)
					close(s.conn.send)
					if len(connections) == 0 {
						delete(Hub.rooms, s.room)
					}
				}
			}
		case m := <-Hub.broadcast:
			connections := Hub.rooms[m.room]
			for c := range connections {
				select {
				case c.send <- m.data:
				default:
					close(c.send)
					delete(connections, c)
					if len(connections) == 0 {
						delete(Hub.rooms, m.room)
					}
				}
			}
		}
	}
}
