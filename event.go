package whatsauth

import (
	"database/sql"
	"log"

	"github.com/whatsauth/watoken"
)

func EventReadSocket(roomId string, PublicKey string, usertables []LoginInfo, db *sql.DB) {
	if roomId[0:1] == "v" {
		phonenumber := watoken.DecodeGetId(PublicKey, roomId)
		if phonenumber != "" {
			infologin := GetLoginInfofromPhoneNumber(phonenumber, usertables, db)
			infologin.Uuid = roomId
			log.Println(infologin)
			SendStructTo(roomId, infologin)
		}
	}
}
