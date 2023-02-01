package whatsauth

import "github.com/aiteung/atmodel"

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

func GenerateButtonMessageCustom(header string, content string, footer string, button []atmodel.WaButton) (btnmsg atmodel.ButtonsMessage) {
	btnmsg.Message.HeaderText = header
	btnmsg.Message.ContentText = content
	btnmsg.Message.FooterText = footer
	btnmsg.Buttons = button
	return
}
