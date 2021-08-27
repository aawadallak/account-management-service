package services

import "golang.org/x/crypto/bcrypt"

type Encrypt struct {
}

func (e *Encrypt) EncryptPassword(password string) (string, error) {

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func (e *Encrypt) CompareEncryptPassword(password, hash string) (bool, error) {

	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	if err != nil {
		return false, err
	}

	return true, nil
}
