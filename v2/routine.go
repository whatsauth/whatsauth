package whatsauth

import (
	"github.com/JPratama7/util/convert"
	"github.com/goccy/go-json"
	"github.com/gofiber/contrib/websocket"
	"github.com/puzpuzpuz/xsync"
	"log"
)

func NewHub(
	reg chan *Client,
	send chan *Message,
	unreg chan string,
	clients *xsync.MapOf[string, *websocket.Conn]) *Hub {
	return &Hub{
		register:   reg,
		send:       send,
		unregister: unreg,
		clients:    clients,
	}
}

type Hub struct {
	register   chan *Client
	send       chan *Message
	unregister chan string
	clients    *xsync.MapOf[string, *websocket.Conn]
}

func (h Hub) Run() {
	for {
		select {
		case connection := <-h.register:
			h.clients.Store(connection.Id, connection.Conn)
			log.Println("connection registered")
			log.Println(connection)
			poolClient.Put(connection)

		case message := <-h.send:
			connection, ok := h.clients.Load(message.Id)
			if !ok {
				log.Println("connection not found 1")
				poolMessage.Put(message)
				continue
			}
			err := connection.WriteMessage(websocket.TextMessage, []byte(message.Message))
			log.Printf("message send to : %+v\n", message)
			if err != nil {
				log.Println(err)
			}
			poolMessage.Put(message)

		case connection := <-h.unregister:
			// Remove the client from the hub
			c, ok := h.clients.LoadAndDelete(connection)
			if !ok {
				log.Println("connection not found 2")
				continue
			}
			c.Close()
			c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			log.Printf("connection unregistered %s\n", connection)
		}
	}
}

func (h Hub) SendMessageTo(ID string, msg string) (res bool) {
	m := poolMessage.Get()
	m.Id = ID
	m.Message = msg

	if _, ok := h.clients.Load(ID); !ok {
		res = false
	} else {
		h.send <- m
		res = true
	}
	return
}

func (h Hub) SendStructTo(ID string, data any) (res bool) {
	b, err := json.Marshal(data)
	if err != nil {
		return
	}

	return h.SendMessageTo(ID, convert.SafeString(b))
}

func (h Hub) Register(client *Client) {
	log.Printf("Register ID: %s\n", client.Id)
	h.register <- client
}

func (h Hub) UnRegister(ID string) {
	log.Printf("UnRegister ID: %s\n", ID)
	h.unregister <- ID
}
