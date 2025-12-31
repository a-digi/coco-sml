package encrypt

import (
	"golang.org/x/crypto/bcrypt"
)

func BcryptHashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func BcryptVerifyPassword(password, hashed string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password)) == nil
}

type BcryptEncode struct{}

func (b *BcryptEncode) PasswordEncrypt(password string) (string, error) {
	return BcryptHashPassword(password)
}

type BcryptVerify struct{}

func (b *BcryptVerify) PasswordVerify(password string, hash string) bool {
	return BcryptVerifyPassword(password, hash)
}

type BcryptManager struct {
	Encoder  PasswordEncrypt
	Verifier PasswordVerifier
}

func CreateBcryptManager() *BcryptManager {
	return &BcryptManager{
		Encoder:  &BcryptEncode{},
		Verifier: &BcryptVerify{},
	}
}
