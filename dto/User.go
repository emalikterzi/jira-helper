package dto

type User struct {
	token string
}

func (u *User) StoreToken(token string) {
	u.token = token;
}

func (u *User) GetToken() (string) {

	return u.token;
}

func (u *User) ClearToken() {
	u.token = "";
}