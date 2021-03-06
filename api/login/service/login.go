package service

type LoginService interface {
	LogInUser(email string, password string) bool
}

type LoginInformation struct {
	email    string
	password string
}

func StaticLoginService() LoginService {
	return &LoginInformation{
		email:    "evsharonov@gmail.com",
		password: "qwerty",
	}
}

func (info *LoginInformation) LogInUser(email, password string) bool {
	return info.email == email && info.password == password
}
