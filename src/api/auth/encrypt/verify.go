package encrypt

type PasswordVerifier interface {
	PasswordVerify(password string, hash string) bool
}
