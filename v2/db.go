package whatsauth

import (
	"context"
	"github.com/JPratama7/util/hunch"
	"github.com/whatsauth/watoken"
)

func GetLoginInfofromPhoneNumber(phonenumber string, usertables []Queries) (response LoginInfo, err error) {
	for _, table := range usertables {
		response, err = table.GetLoginInfoByPhone(phonenumber)
		if err != nil {
			return
		}
	}

	return
}

func GetUsernamefromPhonenumber(phone_number string, usertables []Queries) (username string, err error) {
	for _, table := range usertables {
		username, err = table.GetUsernameByPhone(phone_number)
		if err == nil {
			return
		}
	}

	return
}
func CheckIfUsernameExist(uname, phone_number string, usertables []Queries) (username string, err error) {
	for _, table := range usertables {
		username, err = table.GetUsernameByUnamePhone(uname, phone_number)
		if err == nil {
			return
		}
	}
	return
}

func GetRolesByPhonenumber(
	username,
	phoneNumber string,
	usertables []Queries,
) (loginInfo LoginInfo, err error) {
	for _, table := range usertables {
		loginInfo, err = table.GetLoginInfoByPhoneUname(username, phoneNumber)
		if err == nil {
			break
		}
	}

	pass := watoken.RandomLowerCaseString(21)

	listF := make([]hunch.Executable[*struct{}], 0, len(usertables))

	for _, table := range usertables {
		listF = append(listF, func(ctx context.Context) (*struct{}, error) {
			_, err = table.UpdatePasswordByUsername(loginInfo.Username, pass)
			return nil, err
		})
	}

	_ = hunch.ThrowMut[*struct{}](context.Background(), listF...)

	for _, table := range usertables {
		loginInfo.Userid, err = table.GetUserIdByUsername(loginInfo.Username)
		if err != nil {
			continue
		}
	}

	return
}

func GetListUsernamefromPhonenumber(phone_number string, usertables []Queries) (usernames []string) {
	for _, table := range usertables {
		uname, err := table.GetUsernameByPhone(phone_number)
		if err != nil {
			continue
		}
		if uname == "" {
			continue
		}
		usernames = append(usernames, uname)
	}
	return
}
