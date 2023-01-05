package whatsauth

func SendMessageTo(userID string, msg string) error {
	m := message{[]byte(msg), userID}
	Hub.broadcast <- m
	return nil
}
