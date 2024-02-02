package whatsauth

import (
	"context"
	"github.com/JPratama7/util/hunch"
)

// Wrapper is a function that wraps the Queries.GetUsernameByPhone method
func UsernameWrapper(table Queries, phoneNum string) func(ctx context.Context) (string, error) {
	return func(ctx context.Context) (string, error) {
		data, err := table.GetUsernameByPhone(phoneNum)
		return data, err
	}
}

// Wrapper is a function that wraps the Queries.GetUsernamesByPhone method
func UsernamesWrapper(table Queries, phoneNum string) func(ctx context.Context) ([]string, error) {
	return func(ctx context.Context) ([]string, error) {
		data, err := table.GetUsernamesByPhone(phoneNum)
		return data, err
	}
}

// Wrapper is a function that wraps the Queries.GetUsernameByPhone method
func UpdatePasswordWrapper(table Queries, userName, password string) func(ctx context.Context) (*struct{}, error) {
	return func(ctx context.Context) (*struct{}, error) {
		var locName, locPass = userName, password
		_, err := table.UpdatePasswordByUsername(locName, locPass)
		return nil, err
	}
}

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

	listF := make([]hunch.Executable[*struct{}], 0, len(usertables))

	for _, table := range usertables {
		elems := UpdatePasswordWrapper(table, loginInfo.Username, loginInfo.Password)
		listF = append(listF, elems)
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

	listF := make([]hunch.Executable[[]string], 0, len(usertables))
	for _, table := range usertables {
		elems := UsernamesWrapper(table, phone_number)
		listF = append(listF, elems)
	}
	res, _ := hunch.AllMut(context.Background(), true, listF...)

	for _, val := range res {
		if len(val) > 0 {
			usernames = append(usernames, val...)
		}
	}

	return
}
