package whatsauth

import (
	"fmt"
	"github.com/JPratama7/util/convert"
	"github.com/aiteung/atmodel"
	"github.com/gofiber/contrib/websocket"
	"github.com/whatsauth/watoken"
	"log"
)

func Websocket(hub *Hub, publicKey string, usertables []Queries) func(*websocket.Conn) {
	return func(conn *websocket.Conn) {
		cl := poolClient.Get()
		cl.Conn = conn

		defer func() {
			hub.UnRegister(cl.Id)
			conn.Close()
		}()

		mt, id, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Error reading message: %v", err)
			return
		}
		if mt != websocket.TextMessage {
			conn.WriteMessage(websocket.CloseMessage, convert.UnsafeBytes("not accepted message type"))
			return
		}
		if len(id) < 1 {
			conn.WriteMessage(websocket.CloseMessage, convert.UnsafeBytes("not accepted message type"))
			return
		}

		cl.Id = convert.SafeString(id)
		hub.Register(cl)

		EventReadSocket(hub, cl.Id, publicKey, usertables)

		for {
			i, m, e := conn.ReadMessage()
			fmt.Printf("message: %s\n", e)
			if e != nil {
				break
			}

			e = conn.WriteMessage(i, m)
			if e != nil {
				break
			}
		}
	}
}

func RunWithUsernames(hub *Hub, req WhatsauthRequest, PrivateKey string, usertables []Queries) atmodel.NotifButton {
	header := "WhatsAuth Single Sign On"
	var content string
	footer := fmt.Sprintf("Aplikasi : %v", watoken.GetAppSubDomain(req.Uuid))
	delay := req.Delay
	usernames := make([]string, 0)
	if temp, _ := GetUsernamefromPhonenumber(req.Phonenumber, usertables); temp != "" {
		usernames = GetListUsernamefromPhonenumber(req.Phonenumber, usertables)
		content = fmt.Sprintf("%v detik menunggu kakak mengirim pesan diatas.\nSelanjutnya kakak *klik* di bawah ini ya untuk memilih username kakak.", delay)
	} else {
		content = fmt.Sprintf("Hai kak , Nomor whatsapp ini *tidak terdaftar* di sistem kami, silahkan silahkan gunakan nomor yang terdftar ya kak. Waktu scan %v detik.", delay)
	}

	usernames = DuplicateRemover(usernames)

	if len(usernames) == 1 {
		data := WhatsAuthRoles{
			Uuid:        req.Uuid,
			Phonenumber: req.Phonenumber,
			Roles:       usernames[0],
		}
		log.Printf("\nAuto Select Username : %+v\n", data)
		return SelectedRoles(hub, data, PrivateKey, usertables)
	}

	btm := GenerateButtonMessageCustom(req.Uuid, header, content, footer, usernames)
	var notifbtn atmodel.NotifButton
	notifbtn.User = req.Phonenumber
	notifbtn.Server = "s.whatsapp.net"
	notifbtn.Message = btm
	return notifbtn
}

func SelectedRoles(
	hub *Hub,
	req WhatsAuthRoles,
	PrivateKey string,
	usertables []Queries,
) (notifbtn atmodel.NotifButton) {
	content := ""
	footer := fmt.Sprintf("Aplikasi : %v", watoken.GetAppSubDomain(req.Uuid))
	header := ""
	if temp, _ := CheckIfUsernameExist(req.Roles, req.Phonenumber, usertables); temp != "" {
		infologin, _ := GetRolesByPhonenumber(req.Roles, req.Phonenumber, usertables)
		header = fmt.Sprintf("Silahkan masuk sebagai %s", infologin.Username)
		infologin.Uuid = req.Uuid
		infologin.Login, _ = watoken.EncodeWithStruct(req.Phonenumber, &LoginData{
			Username:    infologin.Username,
			Password:    infologin.Password,
			PhoneNumber: infologin.Phone,
		}, PrivateKey)
		status := hub.SendStructTo(req.Uuid, infologin)
		if status {
			content = "Hai kak , login aplikasi *sukses*,\nsilahkan kakak kembali ke aplikasi."
		} else {
			if req.Uuid[0:1] == "m" {
				content = "\n Selanjutnya kakak *tinggal login* ya kak."
				tokenstring, err := watoken.EncodeforSeconds(req.Phonenumber, PrivateKey, 30)
				if err != nil {
					fmt.Println("simpati RunWithUsername : ", err)
				}
				urlakses := watoken.GetAppUrl(req.Uuid) + "?uuid=" + tokenstring
				footer = fmt.Sprintf("Magic Link : %v", urlakses)
			} else {
				content = "Maaf kak *login gagal*.\nKemungkinan qr code tidak valid atau qr code nya sudah expire kak, silahkan scan ulang kembali ya kak. Semangat kak."
			}
		}
	} else {
		content = "Hai kak , Nomor whatsapp ini *tidak terdaftar* di sistem kami, silahkan silahkan gunakan nomor yang terdftar ya kak."
	}
	notifbtn.User = req.Phonenumber
	notifbtn.Server = "s.whatsapp.net"
	notifbtn.Message = GenerateButtonMessage(header, content, footer)
	return

}
