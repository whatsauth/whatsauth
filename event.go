package whatsauth

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

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
	} else if roomId[0:1] == "g" {
		token := strings.SplitN(roomId, ".", 3)
		if len(token) < 3 {
			return
		}

		phonenumber := watoken.DecodeGetId(PublicKey, token[2])
		if phonenumber != "" {
			infologin := GetRolesByPhonenumber(phonenumber, token[1], usertables, db)
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

func KeyGenerator(roomId string, PublicKey string) {
	if roomId[0:1] != "v" {
		return
	}

	phonenumber := watoken.DecodeGetId(PublicKey, roomId)
	if phonenumber != "" {
		infologin := LoginInfo{Phone: phonenumber, Uuid: roomId}
		SendStructTo(roomId, infologin)
	}
}
