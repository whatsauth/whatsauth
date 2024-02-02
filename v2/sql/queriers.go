package sql

import (
	"auth-service/whatsauth"
	"fmt"
	"github.com/whatsauth/watoken"
)

func (q *Queriers) Ping() (err error) {
	q.mut.Lock()
	defer q.mut.Unlock()

	err = q.db.Ping()
	return
}

func (q *Queriers) GetUsernameByPhone(phoneNumber string) (username string, err error) {
	if err = q.Ping(); err != nil {
		return
	}

	q.mut.Lock()
	defer q.mut.Unlock()

	tsql := fmt.Sprintf(getUsername, q.config.Username, q.config.Uuid,
		q.config.Phone, phoneNumber)

	err = q.db.QueryRow(tsql).Scan(&username)
	if username == "" {
		err = fmt.Errorf("no data found")
		return
	}
	return
}
func (q *Queriers) GetUsernamesByPhone(phoneNumber string) (username []string, err error) {
	if err = q.Ping(); err != nil {
		return
	}

	q.mut.Lock()
	defer q.mut.Unlock()

	tsql := fmt.Sprintf(getUsername, q.config.Username, q.config.Uuid,
		q.config.Phone, phoneNumber)
	fmt.Printf("%s\n\n", tsql)

	cur, err := q.db.Query(tsql)
	if err != nil {
		return
	}
	defer cur.Close()

	for cur.Next() {
		var uname string
		err = cur.Scan(&uname)
		if err != nil {
			return
		}
		username = append(username, uname)
	}

	if len(username) == 0 {
		err = fmt.Errorf("no data found")
		return
	}
	return
}

func (q *Queriers) GetUserIdByUsername(username string) (userId string, err error) {
	if q.Ping() != nil {
		return
	}

	q.mut.Lock()
	defer q.mut.Unlock()

	tsql := fmt.Sprintf(getUsername, q.config.Userid, q.config.Uuid,
		q.config.Username, username)
	err = q.db.QueryRow(tsql).Scan(&userId)
	return
}

func (q *Queriers) GetUsernameByUnamePhone(uname, phoneNumber string) (username string, err error) {
	if q.Ping() != nil {
		return
	}

	q.mut.Lock()
	defer q.mut.Unlock()

	tsql := fmt.Sprintf(checkUsernameWithPhone, q.config.Username, q.config.Uuid,
		q.config.Phone, phoneNumber, q.config.Username, uname)
	err = q.db.QueryRow(tsql).Scan(&username)
	return
}

func (q *Queriers) UpdatePasswordByUsername(username string, pass string) (password string, err error) {
	password = pass
	if password == "" {
		password = watoken.RandomLowerCaseString(21)
	}

	var hashpass string
	switch q.config.Login {
	case "md5":
		hashpass = watoken.GetMD5Hash(password)
	case "2md5":
		hashpass = watoken.GetMD5Hash(watoken.GetMD5Hash(password))
	case "bcrypt":
		hashpass = watoken.GetBcryptHash(password)
	default:
		hashpass = password
	}

	q.mut.Lock()
	defer q.mut.Unlock()

	tsql := fmt.Sprintf(updatePassword, q.config.Uuid,
		q.config.Password, hashpass,
		q.config.Username, username)
	res, err := q.db.Exec(tsql)
	if af, _ := res.RowsAffected(); af == 0 {
		err = fmt.Errorf("no rows affected")
	}

	return
}

func (q *Queriers) GetLoginInfoByPhone(phoneNumber string) (data whatsauth.LoginInfo, err error) {
	data.Username, err = q.GetUsernameByPhone(phoneNumber)
	if err != nil {
		return
	}

	data.Password, err = q.UpdatePasswordByUsername(data.Username, "")
	if err != nil {
		return
	}

	data.Userid, err = q.GetUserIdByUsername(data.Username)
	return
}

func (q *Queriers) GetLoginInfoByPhoneUname(uname, phoneNumber string) (data whatsauth.LoginInfo, err error) {
	data.Username, err = q.GetUsernameByUnamePhone(uname, phoneNumber)
	if err != nil {
		return
	}

	data.Password, err = q.UpdatePasswordByUsername(data.Username, "")
	if err != nil {
		return
	}

	data.Userid, err = q.GetUserIdByUsername(data.Username)
	return
}
