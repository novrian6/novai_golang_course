package models

type user struct {
	Username string `json:"username"`
	Password string `json:"-"`
}

var userList = []user{
	user{Username: "user1", Password: "pass1"},

}

func IsUserValid(username, password string) bool {
	for _, u:= range userList {
		if u.Username == username && u.Password == password   {

			return true
		}
	}
	return false
}