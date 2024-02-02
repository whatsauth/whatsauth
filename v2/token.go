package whatsauth

import (
	"github.com/whatsauth/watoken"
	"log"
)

func EventReadSocket(ru *Hub, roomId string, PublicKey string, usertables []Queries) {
	if roomId[:1] != "v" {
		return
	}

	infologin := LoginInfo{}

	defer func() {
		infologin.Uuid = roomId
		infologin.Login = roomId
		ru.SendStructTo(roomId, infologin)
	}()

STARTFUNC:
	phonenumber, err := watoken.DecodeWithStruct[LoginData](PublicKey, roomId)
	if err != nil {
		log.Printf("Error decode %v\n", err)
		return
	}

	if phonenumber.Data.Password != "" && phonenumber.Data.Username != "" {
		infologin.Username = phonenumber.Data.Username
		infologin.Password = phonenumber.Data.Password
		if err != nil {
			log.Printf("Error GetLoginInfofromUsername %v\n", err)
			return
		}
	}

	if phonenumber.Data.Username != "" {
		infologin, err = GetRolesByPhonenumber(phonenumber.Data.Username, phonenumber.Id, usertables)
		if err != nil {
			log.Printf("Error GetRolesByPhonenumber %v\n", err)
			return
		}

		log.Println("Info Login EventReadSocket Username ", infologin)
	}

	if phonenumber.Id != "" {
		infologin, err = GetLoginInfofromPhoneNumber(phonenumber.Id, usertables)
		if err != nil {
			log.Printf("Error GetLoginInfofromPhoneNumber %v\n", err)
			return
		}
		log.Printf("Info Login EventReadSocket PhoneNumber %+v\n", infologin)
	}

	if infologin.Username == "" {
		goto STARTFUNC
	}

}
