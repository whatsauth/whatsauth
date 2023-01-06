package whatsauth

func SendMessageTo(userID string, msg string) (res string) {
	m := message{[]byte(msg), userID}
	connections := Hub.rooms[userID]
	if connections == nil {
		res = "notfound"
	} else {
		Hub.broadcast <- m
		res = "sent"
	}
	return res
}
