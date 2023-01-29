package whatsauth

import (
	"database/sql"
	"fmt"

	"github.com/whatsauth/wasocket"
	"github.com/whatsauth/watoken"

	"github.com/aiteung/atmodel"
)

func RunWS(roomId string, PublicKey string, usertables []LoginInfo, db *sql.DB) {
	if roomId[0:1] == "v" {
		phonenumber := watoken.DecodeGetId(PublicKey, roomId)
		if phonenumber != "" {
			infologin := GetLoginInfofromPhoneNumber(phonenumber, usertables, db)
			infologin.Uuid = roomId
			fmt.Println(infologin)
			wasocket.SendStructTo(roomId, infologin)
		}
	}
}

func RunModule(req WhatsauthRequest, PrivateKey string, usertables []LoginInfo, db *sql.DB) atmodel.NotifButton {
	header := "WhatsAuth Single Sign On"
	var content string
	footer := fmt.Sprintf("Aplikasi : %v", watoken.GetAppSubDomain(req.Uuid))
	delay := req.Delay
	if GetUsernamefromPhonenumber(req.Phonenumber, usertables, db) != "" {
		infologin := GetLoginInfofromPhoneNumber(req.Phonenumber, usertables, db)
		infologin.Uuid = req.Uuid
		infologin.Login, _ = watoken.Encode(infologin.Username, PrivateKey)
		fmt.Println(infologin)
		status := wasocket.SendStructTo(req.Uuid, infologin)
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

func GenerateButtonMessage(header string, content string, footer string) (btnmsg atmodel.ButtonsMessage) {
	btnmsg.Message.HeaderText = header
	btnmsg.Message.ContentText = content
	btnmsg.Message.FooterText = footer
	btnmsg.Buttons = []atmodel.WaButton{{
		ButtonId:    "whatsauth|challange1",
		DisplayText: "Sama Sama",
	},
		{
			ButtonId:    "whatsauth|challange3",
			DisplayText: "Sawangsulna",
		},
		{
			ButtonId:    "whatsauth|challange2",
			DisplayText: "Mangga",
		},
	}
	return btnmsg
}
