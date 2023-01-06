package whatsauth

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
