package whatsauth

import (
	"fmt"
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

func HasRoleKeyword(Info *types.MessageInfo, Message *waProto.Message) (whmsg bool) {
	if Message.ExtendedTextMessage != nil && Info.Chat.Server == "s.whatsapp.net" {
		if strings.Contains(*Message.ExtendedTextMessage.Text, Keyword+".") {
			if Message.ExtendedTextMessage.ContextInfo != nil {
				if *Message.ExtendedTextMessage.ContextInfo.EntryPointConversionSource == "click_to_chat_link" {
					whmsg = true
				}
			}
		}
	}
	return
}

func RunModule(waclient *whatsmeow.Client, Info *types.MessageInfo, Message *waProto.Message, urlwauthreq string, prefixurlapiwa string) {
	var wareq WhatsauthRequest
	wareq.Uuid = strings.Replace(*Message.ExtendedTextMessage.Text, Keyword, "", 1)
	wareq.Phonenumber = Info.Sender.User
	wareq.Delay = *Message.ExtendedTextMessage.ContextInfo.EntryPointConversionDelaySeconds
	fmt.Println(wareq)
	ntfbtn, _ := atapi.PostStruct[atmessage.NotifButton](wareq, urlwauthreq)
	fmt.Println(ntfbtn)
	btm := ntfbtn.Message
	atmessage.SendMessage(ButtonMessageToMessage(btm, prefixurlapiwa), Info.Sender, waclient)
	resp, err := atmessage.SendButtonMessage(btm, Info.Sender, waclient)
	fmt.Println(resp)
	fmt.Println(err)
}

func HandlerWhatsauth(waclient *whatsmeow.Client, Info *types.MessageInfo, Message *waProto.Message, urlwauthreq string, urlwauthrole string, prefixurlapiwa string) {
	if Message.ButtonsResponseMessage != nil {
		ButtonMessageWhatsauth(waclient, Info, Message, urlwauthrole)
	} else if HasRoleKeyword(Info, Message) {
		MessageWhatsauth(waclient, Info, Message, urlwauthrole)
	} else {
		RunModule(waclient, Info, Message, urlwauthreq, prefixurlapiwa)
	}
}

func MessageWhatsauth(waclient *whatsmeow.Client, Info *types.MessageInfo, Message *waProto.Message, urlwauthrole string) {
	wareq := new(WhatsAuthRoles)
	uuidSplit := strings.SplitN(*Message.ExtendedTextMessage.Text, ".", 3)
	if len(uuidSplit) < 2 {
		atmessage.SendMessage("Error, len split UUID from ButtonID not enough.", Info.Chat, waclient)

	}
	wareq.Uuid = uuidSplit[2]
	wareq.Roles = uuidSplit[1]
	wareq.Phonenumber = Info.Sender.User
	fmt.Println(wareq)
	ntfbtn, _ := atapi.PostStruct[atmessage.NotifButton](wareq, urlwauthrole)
	fmt.Println(ntfbtn)
	btm := ntfbtn.Message
	if btm.Message.HeaderText != "" {
		atmessage.SendMessage(ButtonMessageToMessage(btm, ""), Info.Sender, waclient)
		resp, err := atmessage.SendButtonMessage(btm, Info.Sender, waclient)
		fmt.Println(resp)
		fmt.Println(err)
	} else {
		fmt.Println("struct kembalian kosong")
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
	fmt.Println(wareq)
	ntfbtn, _ := atapi.PostStruct[atmessage.NotifButton](wareq, urlwauthrole)
	fmt.Println(ntfbtn)
	btm := ntfbtn.Message
	if btm.Message.HeaderText != "" {
		resp, err := atmessage.SendButtonMessage(btm, Info.Sender, waclient)
		fmt.Println(resp)
		fmt.Println(err)
	} else {
		fmt.Println("struct kembalian kosong")
	}
}

func FilterWhatsauthButton(Info *types.MessageInfo, Message *waProto.Message) (filter bool) {
	if Message.ButtonsResponseMessage != nil && Info.Chat.Server == "s.whatsapp.net" {
		filter = strings.Contains(*Message.ButtonsResponseMessage.SelectedButtonId, Keyword)
	}
	return
}
