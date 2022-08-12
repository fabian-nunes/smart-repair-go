package auth

type User struct {
	Username string
	Password string
}

func newUser(username string, password string) User {
	return User{
		Username: username,
		Password: password,
	}
}
