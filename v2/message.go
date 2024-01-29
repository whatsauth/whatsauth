package whatsauth

import (
	"strconv"
	"strings"

	"github.com/aiteung/atmessage"
	"github.com/aiteung/atmodel"
)

func GenerateButtonMessage(header string, content string, footer string) (btnmsg atmodel.ButtonsMessage) {
	btnmsg.Message.HeaderText = header
	btnmsg.Message.ContentText = content
	btnmsg.Message.FooterText = footer
	btnmsg.Buttons = []atmodel.WaButton{}
	return btnmsg
}

func GenerateButtonMessageCustom(
	uuid string,
	header string,
	content string,
	footer string,
	button []string,
) (btnmsg atmodel.ButtonsMessage) {
	if len(button) == 0 {
		btnmsg = GenerateButtonMessage(header, content, footer)
		return
	}
	keys := make(map[string]bool, len(button))
	uniq := make([]string, 0, len(button))
	for _, entry := range button {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			uniq = append(uniq, entry)
		}
	}

	btnmsg.Message.HeaderText = header
	btnmsg.Message.ContentText = content
	btnmsg.Message.FooterText = footer
	butt := make([]atmodel.WaButton, 0, len(uniq))
	for _, v := range uniq {
		butt = append(butt, atmodel.WaButton{
			ButtonId:    "wh4t5auth0." + v + "." + uuid,
			DisplayText: strings.ToTitle(v),
		})
	}
	btnmsg.Buttons = butt
	return
}

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
