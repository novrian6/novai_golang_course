package models

type User struct {
	IDUsers string ` json:"idusers"`
	Name string `gorm:"type: varchar(255)" json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Status string `json:"status"`
	Email  string `json:"email"`
	Token string `json:"token,omitempty""`
}

var userList = []User{
	User{Username: "user1", Password: "pass1"},

}

func IsUserValid(username, password string) bool {
	for _, u:= range userList {
		if u.Username == username && u.Password == password   {

			return true
		}
	}
	return false
}

