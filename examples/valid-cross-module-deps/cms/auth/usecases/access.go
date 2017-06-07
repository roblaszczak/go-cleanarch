package usecases

// LoginAccessChecker is used by CheckAccess in infrastructure layer.
func LoginAccessChecker(username string) bool {
	return username == "admin"
}
