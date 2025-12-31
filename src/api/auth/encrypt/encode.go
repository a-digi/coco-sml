package encrypt

type PasswordEncode interface {
	PasswordEncode(password string) (string, error)
}

