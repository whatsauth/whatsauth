package whatsauth

import (
	"github.com/JPratama7/util/sync"
	sy "sync"
)

var poolClient *sync.Pool[*Client]

var poolMessage *sync.Pool[*Message]

var once sy.Once

func init() {
	once.Do(func() {
		poolMessage = sync.NewPool(func() *Message {
			return new(Message)
		})

		poolClient = sync.NewPool(func() *Client {
			return new(Client)
		})
	})
}
