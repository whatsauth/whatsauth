package whatsauth

import (
	"encoding/json"
	"fmt"
)

func SendMessageTo(ID string, msg string) (res bool) {
	m := message{[]byte(msg), ID}
	if Hub.rooms[ID] == nil {
		res = false
	} else {
		Hub.broadcast <- m
		res = true
	}
	return res
}

func SendStructTo(ID string, strc interface{}) (res bool) {
	b, err := json.Marshal(strc)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
	return SendMessageTo(ID, string(b))
}
