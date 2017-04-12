package domain

func NewUser(username string) User {
	return User{username}
}

type User struct {
	username string
}

func (u User) Username() string {
	return u.username
}
