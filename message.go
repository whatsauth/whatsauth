package whatsauth

import (
	"github.com/aiteung/atmodel"
	"strings"
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
			ButtonId:    v,
			DisplayText: strings.ToTitle(v),
		})
	}
	btnmsg.Buttons = butt
	return
}
