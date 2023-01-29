package whatsauth

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gofiber/websocket/v2"
	"github.com/whatsauth/watoken"

	"github.com/aiteung/atmodel"
)

func RunHub() { // Call this function on your main function before run fiber
	for {
		select {
		case connection := <-Register:
			Clients[connection.Id] = connection.Conn
			log.Println("connection registered")
			log.Println(connection)

		case message := <-SendMesssage:
			log.Println("message received:", message)
			connection := Clients[message.Id]
			err := connection.WriteMessage(websocket.TextMessage, []byte(message.Message))
			if err != nil {
				log.Println(err)
			}

		case connection := <-Unregister:
			// Remove the client from the hub
			delete(Clients, connection)

			log.Println("connection unregistered")
			log.Println(connection)
		}
	}
}

func RunSocket(c *websocket.Conn, PublicKey string, usertables []LoginInfo, db *sql.DB) (Id string) { // call this function after declare URL routes
	var s Client
	// When the function returns, unregister the client and close the connection
	defer func() {
		Unregister <- s.Id
		c.Close()
	}()
	messageType, message, err := c.ReadMessage()
	if err != nil {
		if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
			log.Println("read error:", err)
		}
		return // Calls the deferred function, i.e. closes the connection on error
	}
	Id = string(message)
	if messageType == websocket.TextMessage {
		// Get the received message
		// Register the client
		s = Client{
			Id:   Id,
			Conn: c,
		}
		Register <- s
		EventReadSocket(Id, PublicKey, usertables, db)
		for {
			messageType, message, err := s.Conn.ReadMessage()
			if err != nil {
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
					log.Println("read error:", err)
				}
				return // Calls the deferred function, i.e. closes the connection on error
			}

			if messageType == websocket.TextMessage {
				// log the received message
				log.Println(string(message))
			} else {
				log.Println("websocket message received of type", messageType)
			}
		}
	} else {
		log.Println("websocket message received of type", messageType)
	}
	return

}

func RunModule(req WhatsauthRequest, PrivateKey string, usertables []LoginInfo, db *sql.DB) atmodel.NotifButton {
	header := "WhatsAuth Single Sign On"
	var content string
	footer := fmt.Sprintf("Aplikasi : %v", watoken.GetAppSubDomain(req.Uuid))
	delay := req.Delay
	if GetUsernamefromPhonenumber(req.Phonenumber, usertables, db) != "" {
		infologin := GetLoginInfofromPhoneNumber(req.Phonenumber, usertables, db)
		infologin.Uuid = req.Uuid
		infologin.Login, _ = watoken.Encode(req.Phonenumber, PrivateKey)
		fmt.Println(infologin)
		status := SendStructTo(req.Uuid, infologin)
		if status {
			content = fmt.Sprintf("Hai kak , login aplikasi *sukses*,\nsilahkan kakak kembali ke aplikasi.\nLama kakak kirim pesan di atas : %v detik.", delay)
		} else {
			if req.Uuid[0:1] == "m" {
				content = fmt.Sprintf("%v detik menunggu kakak mengirim pesan diatas.\nSelanjutnya kakak *buka Magic Link* di bawah ini ya kak, link berlaku selama 30 detik.", delay)
				tokenstring, err := watoken.EncodeforSeconds(req.Phonenumber, PrivateKey, 30)
				if err != nil {
					fmt.Println("simpati RunModule : ", err)
				}
				urlakses := watoken.GetAppUrl(req.Uuid) + "?uuid=" + tokenstring
				footer = fmt.Sprintf("Magic Link : %v", urlakses)
			} else {
				content = fmt.Sprintf("Maaf kak *login gagal*.\nKemungkinan qr code tidak valid atau qr code nya sudah expire kak, silahkan scan ulang kembali ya kak.\nKakak butuh waktu %v detik untuk mengirim pesan diatas. Semoga selanjutnya bisa lebih cekatan ya kak. Semangat kak.", delay)
			}

		}
	} else {
		content = fmt.Sprintf("Hai kak , Nomor whatsapp ini *tidak terdaftar* di sistem kami, silahkan silahkan gunakan nomor yang terdftar ya kak. Waktu scan %v detik.", delay)
	}
	btm := GenerateButtonMessage(header, content, footer)
	var notifbtn atmodel.NotifButton
	notifbtn.User = req.Phonenumber
	notifbtn.Server = "s.whatsapp.net"
	notifbtn.Message = btm
	return notifbtn
}
