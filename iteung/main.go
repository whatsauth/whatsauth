package iteung

import (
	"fmt"
	"strings"

	"github.com/aiteung/atapi"
	"github.com/aiteung/atmessage"
	"github.com/aiteung/module/model"
)

func MessageWhatsauth(Pesan model.IteungMessage, urlwauthrole string) (msg string) {
	wareq := new(WhatsAuthRoles)
	uuidSplit := strings.SplitN(Pesan.Message, ".", 3)
	if len(uuidSplit) < 2 {
		msg = "Error, len split UUID from ButtonID not enough."
	}
	wareq.Uuid = uuidSplit[2]
	wareq.Roles = uuidSplit[1]
	wareq.Phonenumber = Pesan.Phone_number
	fmt.Println(wareq)
	ntfbtn := atapi.PostStruct[atmessage.NotifButton](wareq, urlwauthrole)
	fmt.Println(ntfbtn)
	btm := ntfbtn.Message
	if btm.Message.HeaderText != "" {
		msg = ButtonMessageToMessage(btm, "")
	} else {
		msg = "struct kembalian kosong"
	}
	return
}

func RunModule(Pesan model.IteungMessage, Keyword string, urlwauthreq string, prefixurlapiwa string) (msg string) {
	var wareq WhatsauthRequest
	wareq.Uuid = strings.Replace(Pesan.Message, Keyword, "", 1)
	wareq.Phonenumber = Pesan.Phone_number
	wareq.Delay = Pesan.From_link_delay
	fmt.Println(wareq)
	ntfbtn := atapi.PostStruct[atmessage.NotifButton](wareq, urlwauthreq)
	fmt.Println(ntfbtn)
	btm := ntfbtn.Message
	msg = ButtonMessageToMessage(btm, prefixurlapiwa)
	return
}
