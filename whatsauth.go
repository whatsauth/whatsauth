package whatsauth

import (
	"strings"

	"github.com/aiteung/atapi"
	"github.com/aiteung/atmessage"
	"go.mau.fi/whatsmeow"
	waProto "go.mau.fi/whatsmeow/binary/proto"
	"go.mau.fi/whatsmeow/types"
)

const Keyword string = "wh4t5auth0"

func HasKeyword(Info *types.MessageInfo, Message *waProto.Message) (whmsg bool) {
	if Message.ExtendedTextMessage != nil && Info.Chat.Server == "s.whatsapp.net" {
		if strings.Contains(*Message.ExtendedTextMessage.Text, Keyword) {
			if Message.ExtendedTextMessage.ContextInfo != nil {
				if *Message.ExtendedTextMessage.ContextInfo.EntryPointConversionSource == "click_to_chat_link" {
					whmsg = true
				}
			}
		}
	}
	return
}

func RunModule(waclient *whatsmeow.Client, Info *types.MessageInfo, Message *waProto.Message, urlwauthreq string) {
	var wareq WhatsauthRequest
	wareq.Uuid = strings.Replace(*Message.ExtendedTextMessage.Text, Keyword, "", 1)
	wareq.Phonenumber = Info.Sender.User
	wareq.Delay = *Message.ExtendedTextMessage.ContextInfo.EntryPointConversionDelaySeconds
	ntfbtn := atapi.PostStruct[atmessage.NotifButton](wareq, urlwauthreq)
	btm := ntfbtn.Message
	atmessage.SendButtonMessage(btm, Info.Sender, waclient)
}

func HandlerWhatsauth(waclient *whatsmeow.Client, Info *types.MessageInfo, Message *waProto.Message, urlwauthreq string, urlwauthrole string) {
	if Message.ButtonsResponseMessage != nil {
		ButtonMessageWhatsauth(waclient, Info, Message, urlwauthrole)
	} else {
		RunModule(waclient, Info, Message, urlwauthreq)
	}
}

func ButtonMessageWhatsauth(waclient *whatsmeow.Client, Info *types.MessageInfo, Message *waProto.Message, urlwauthrole string) {
	wareq := new(WhatsAuthRoles)
	uuidSplit := strings.SplitN(*Message.ButtonsResponseMessage.SelectedButtonId, ".", 3)
	if len(uuidSplit) < 2 {
		atmessage.SendMessage("Error, len split UUID from ButtonID not enough.", Info.Chat, waclient)

	}
	wareq.Uuid = uuidSplit[2]
	wareq.Roles = uuidSplit[1]
	wareq.Phonenumber = Info.Sender.User
	ntfbtn := atapi.PostStruct[atmessage.NotifButton](wareq, urlwauthrole)
	btm := ntfbtn.Message
	atmessage.SendButtonMessage(btm, Info.Sender, waclient)

}

func FilterWhatsauthButton(Info *types.MessageInfo, Message *waProto.Message) (filter bool) {
	if Message.ButtonsResponseMessage != nil && Info.Chat.Server == "s.whatsapp.net" {
		filter = strings.Contains(*Message.ButtonsResponseMessage.SelectedButtonId, Keyword)
	}
	return
}
