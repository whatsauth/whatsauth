package whatsauth

import (
	"github.com/aiteung/atmodel"
	"github.com/whatsauth/watoken"
)

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

func GenerateButtonMessageCustom(header string, content string, footer string, button []string) (btnmsg atmodel.ButtonsMessage) {
	btnmsg.Message.HeaderText = header
	btnmsg.Message.ContentText = content
	btnmsg.Message.FooterText = footer
	butt := make([]atmodel.WaButton, 0, len(button))
	for _, v := range button {
		butt = append(butt, atmodel.WaButton{
			ButtonId:    watoken.RandomLowerCaseString(4),
			DisplayText: v,
		})
	}

	return
}
