package whatsauth

type LoginInfo struct {
	Userid   string `json:"user_id,omitempty" bson:"user_id,omitempty"`
	Username string `json:"user_name,omitempty" bson:"user_name,omitempty"`
	Password string `json:"user_pass,omitempty" bson:"user_pass,omitempty"`
	Phone    string `json:"phone,omitempty" bson:"phone,omitempty"`
	Login    string `json:"login,omitempty" bson:"login,omitempty"`
	Uuid     string `json:"uuid,omitempty" bson:"uuid,omitempty"`
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
