package whatsauth

import (
	"database/sql"
	"fmt"
	"github.com/whatsauth/watoken"
)

func GetLoginInfofromPhoneNumber(phonenumber string, usertables []LoginInfo, db *sql.DB) (response LoginInfo) {
	fmt.Println("phonenumber : " + phonenumber)
	if phonenumber != "" {
		response.Username = GetUsernamefromPhonenumber(phonenumber, usertables, db)
		fmt.Println("username : " + response.Username)
		if response.Username != "" {
			response.Password = UpdatePasswordfromUsername(response.Username, usertables, db)
			fmt.Println("password : " + response.Password)
			if response.Password != "" {
				response.Userid = GetUserIdfromUsername(response.Username, usertables, db)
			}

		}
	}
	return response
}

func GetUsernamefromPhonenumber(phone_number string, usertables []LoginInfo, db *sql.DB) (username string) {
	for _, table := range usertables {
		q := "select %s from %s where %s = '%s'"
		tsql := fmt.Sprintf(q, table.Username, table.Uuid,
			table.Phone, phone_number)
		err := db.QueryRow(tsql).Scan(&username)
		if err == sql.ErrNoRows {
			fmt.Printf("GetUsernamefromPhonenumber, no user in table : %s", table.Uuid)
		} else if err != nil {
			fmt.Printf("GetUsernamefromPhonenumber: %v\n", err)
		}
		if username != "" {
			break
		}
	}
	return
}

func GetListUsernamefromPhonenumber(phone_number string, usertables []LoginInfo, db *sql.DB) (usernames []string) {
	username := ""
	for _, table := range usertables {
		q := "select %s from %s where %s = '%s'"
		tsql := fmt.Sprintf(q, table.Username, table.Uuid,
			table.Phone, phone_number)
		err := db.QueryRow(tsql).Scan(&username)
		if err == sql.ErrNoRows {
			fmt.Printf("GetUsernamefromPhonenumber, no user in table : %s", table.Uuid)
		} else if err != nil {
			fmt.Printf("GetUsernamefromPhonenumber: %v\n", err)
		}
		if username != "" {
			usernames = append(usernames, username)
		}
	}
	return
}

func GetHashPasswordfromUsername(username string, usertables []LoginInfo, db *sql.DB) (hashpassword string) {
	for _, table := range usertables {
		q := "select %s from %s where %s = '%s'"
		tsql := fmt.Sprintf(q, table.Password, table.Uuid,
			table.Username, username)
		err := db.QueryRow(tsql).Scan(&hashpassword)
		if err == sql.ErrNoRows {
			fmt.Printf("GetHashPasswordfromUsername, no user in table : %s", table.Uuid)
		} else if err != nil {
			fmt.Printf("GetHashPasswordfromUsername: %v\n", err)
		}
		if hashpassword != "" {
			break
		}
	}
	return
}

func UpdatePasswordfromUsername(username string, usertables []LoginInfo, db *sql.DB) (newPassword string) {
	newPassword = watoken.RandomLowerCaseString(21)
	for _, table := range usertables {
		var hashpass string
		switch table.Login {
		case "md5":
			hashpass = watoken.GetMD5Hash(newPassword)
		case "2md5":
			hashpass = watoken.GetMD5Hash(watoken.GetMD5Hash(newPassword))
		}
		var temp interface{}
		q := "update %s set %s = '%s' where %s = '%s'"
		tsql := fmt.Sprintf(q, table.Uuid,
			table.Password, hashpass,
			table.Username, username)
		err := db.QueryRow(tsql).Scan(&temp)
		if err == sql.ErrNoRows {
			fmt.Printf("UpdatePasswordfromUsername, Success Update in table : %s", table.Uuid)
		} else if err != nil {
			fmt.Printf("UpdatePasswordfromUsername: %v\n", err)
		}
	}

	return
}

func GetUserIdfromUsername(username string, usertables []LoginInfo, db *sql.DB) (userid string) {
	for _, table := range usertables {
		q := "select %s from %s where %s = '%s'"
		tsql := fmt.Sprintf(q, table.Userid, table.Uuid,
			table.Username, username)
		err := db.QueryRow(tsql).Scan(&userid)
		if err == sql.ErrNoRows {
			fmt.Printf("GetUserIdfromUsername, no user in table : %s", table.Uuid)
		} else if err != nil {
			fmt.Printf("GetUserIdfromUsername: %v\n", err)
		}
		if userid != "" {
			break
		}
	}

	return
}

func GetRolesByPhonenumber(
	phoneNumber string,
	userRoles string,
	usertables []LoginInfo,
	db *sql.DB,
) (loginInfo LoginInfo) {
	loginInfo = GetLoginInfofromPhoneNumber(phoneNumber, usertables, db)
	role := ""
	for _, table := range usertables {
		q := "select %s from %s where %s = '%s' AND %s = '%s'"
		tsql := fmt.Sprintf(q, table.Username, table.Uuid,
			table.Phone, phoneNumber, table.Username, userRoles)
		err := db.QueryRow(tsql).Scan(&role)
		if err == sql.ErrNoRows {
			fmt.Printf("GetRolesByPhonenumber, no user in table : %s \n", table.Uuid)
		} else if err != nil {
			fmt.Printf("GetRolesByPhonenumber: %v\n", err)
		}
	}
	if role != "" {
		loginInfo.Username = role
		loginInfo.Userid = GetUserIdfromUsername(loginInfo.Username, usertables, db)
	}
	return
}
