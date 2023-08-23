package iteung

import (
	"strconv"

	"github.com/aiteung/atmessage"
)

func ButtonMessageToMessage(btm atmessage.ButtonsMessage, prefixurlapiwa string) string {
	judul := "*" + btm.Message.HeaderText + "*\n"
	konten := btm.Message.ContentText + "\n"
	kaki := "_" + btm.Message.FooterText + "_\n\n"
	var listroles string
	for i, role := range btm.Buttons {
		listroles = listroles + strconv.Itoa(i+1) + ". *" + role.DisplayText + "*\n" + prefixurlapiwa + role.ButtonId + "\n"
	}
	return judul + konten + kaki + listroles
}
