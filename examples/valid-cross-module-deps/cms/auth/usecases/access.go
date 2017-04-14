package usecases

func LoginAccessChecker(username string) bool {
	return username == "admin"
}
