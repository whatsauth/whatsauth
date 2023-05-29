package whatsauth

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/whatsauth/watoken"
)

func EventReadSocket(roomId string, PublicKey string, usertables []LoginInfo, db *sql.DB) {
	if roomId[0:1] == "v" {
		phonenumber := watoken.DecodeGetId(PublicKey, roomId)
		if phonenumber != "" {
			infologin := GetLoginInfofromPhoneNumber(phonenumber, usertables, db)
			infologin.Uuid = roomId
			log.Println("Info Login EventReadSocket ", infologin)
			SendStructTo(roomId, infologin)
		} else {
			fmt.Println("EventReadSocket: phonenumber is empty")
		}
	} else {
		log.Println("EventReadSocket: roomId[0:1] != v ", roomId, PublicKey)
	}
}
