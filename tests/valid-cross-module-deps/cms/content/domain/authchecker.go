package domain

type AuthChecker interface {
	CheckAuth(u User) bool
}
