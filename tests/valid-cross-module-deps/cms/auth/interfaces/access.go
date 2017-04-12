package interfaces

func CheckAccess(username string) bool {
	return username == "admin"
}
