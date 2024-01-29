package whatsauth

import "github.com/gofiber/contrib/websocket"

type Client struct {
	Id   string
	Conn *websocket.Conn
} // Register Conn socket with ID

type Message struct {
	Id      string
	Message string
} // To send message to ID

type LoginData struct {
	Username    string `json:"user_name,omitempty" bson:"username,omitempty"`
	Password    string `json:"password,omitempty" bson:"password,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty" bson:"phone_number,omitempty"`
}

type LoginInfo struct {
	Userid   string `json:"user_id,omitempty" bson:"user_id,omitempty"`
	Username string `json:"user_name,omitempty" bson:"user_name,omitempty"`
	Password string `json:"user_pass,omitempty" bson:"user_pass,omitempty"`
	Phone    string `json:"phone,omitempty" bson:"phone,omitempty"`
	Login    string `json:"login,omitempty" bson:"login,omitempty"`
	Uuid     string `json:"uuid,omitempty" bson:"uuid,omitempty"`
}

type Whatsauth struct {
	Uuid string `json:"uuid"`
}

type WhatsauthRequest struct {
	Uuid        string `json:"uuid,omitempty" bson:"uuid,omitempty"`
	Phonenumber string `json:"phonenumber,omitempty" bson:"phonenumber,omitempty"`
	Delay       uint32 `json:"delay,omitempty" bson:"delay,omitempty"`
}

type WhatsauthMessage struct {
	Id      string    `json:"id,omitempty" bson:"id,omitempty"`
	Message LoginInfo `json:"message,omitempty" bson:"message,omitempty"`
}

type WhatsauthStatus struct {
	Status string `json:"status,omitempty" bson:"status,omitempty"`
}

type WhatsAuthRoles struct {
	Uuid        string `json:"uuid,omitempty" bson:"uuid,omitempty"`
	Phonenumber string `json:"phonenumber,omitempty" bson:"phonenumber,omitempty"`
	Roles       string `json:"roles,omitempty" bson:"roles,omitempty"`
}
