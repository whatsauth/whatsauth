package whatsauth

type DataFinderByPhone interface {
	GetUsernameByPhone(phone string) (string, error)
	GetUsernamesByPhone(phone string) ([]string, error)
}

type DataFinderByUsername interface {
	GetUserIdByUsername(username string) (string, error)
}

type UsernameExistChecker interface {
	GetUsernameByUnamePhone(uname, phoneNumber string) (string, error)
}

type PasswordUpdater interface {
	UpdatePasswordByUsername(username string, pass string) (string, error)
}

type LoginFinder interface {
	GetLoginInfoByPhone(phone string) (LoginInfo, error)
	GetLoginInfoByPhoneUname(uname, phone string) (LoginInfo, error)
}

type Queries interface {
	DataFinderByPhone
	DataFinderByUsername
	UsernameExistChecker
	PasswordUpdater
	LoginFinder
}
