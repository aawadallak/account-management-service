package services

import (
	"errors"
	"latest/dto"
	"latest/repository"
)

type UserService struct {
	repository      repository.UserRepository
	redisRepository repository.RedisRepository
}

func (s *UserService) Store(user dto.User) (int64, error) {

	dispatcher := MailDispatcherService{}

	svc := RandomGenerator{}

	user.Code = svc.GenerateRandomString(6)

	err := s.repository.GetByUsername(user.Username)

	if err != nil {
		return 0, err
	}

	service := Encrypt{}

	hash, err := service.EncryptPassword(user.Password)

	if err != nil {
		return 0, err
	}

	w, err := s.repository.Store(user, hash)

	if err != nil {
		return 0, err
	}

	err = dispatcher.SendEmail(user)

	if err != nil {
		return 0, err
	}

	err = s.redisRepository.SendCodeAuth(user.Email, user.Code)

	if err != nil {
		return 0, err
	}

	return w, nil
}

func (s *UserService) LoginWithUserAndPassword(username, password string) (string, error) {

	w, err := s.repository.LoginWithUserAndPassword(username, password)

	if err != nil {
		return "", err
	}

	svc := Encrypt{}

	passwordMatch, err := svc.CompareEncryptPassword(password, w)

	if err != nil {
		return "", err
	}

	if !passwordMatch {
		return "", errors.New("incorrect password or username")
	}

	service := NewJwtService()

	token, err := service.GenJwtToken(w)

	if err != nil {
		return "", err
	}

	return token, nil

}

func (s *UserService) VerifyCode(code string, email string) error {

	_, err := s.repository.GetByUserEmail(email)

	if err != nil {
		return err
	}

	err = s.repository.CheckVerifyStatus(email)

	if err != nil {
		return err
	}

	_, err = s.redisRepository.CheckCode(email, code)

	if err != nil {
		return err
	}

	_, err = s.repository.SetVerifiedTrue(email)

	if err != nil {
		return err
	}

	return nil
}

func (s *UserService) ResendVerifyCode(email string) error {

	dispatcher := MailDispatcherService{}

	err := s.repository.CheckVerifyStatus(email)

	if err != nil {
		return err
	}

	svc := RandomGenerator{}

	code := svc.GenerateRandomString(6)

	err = dispatcher.ReSendEmail(email, code)

	if err != nil {
		return err
	}

	err = s.redisRepository.SendCodeAuth(email, code)

	if err != nil {
		return err
	}

	return nil
}

func (s *UserService) ChangeEmail(user dto.UserUpdateEmail) error {

	dispatcher := MailDispatcherService{}

	svc := RandomGenerator{}

	user.Code = svc.GenerateRandomString(6)

	_, err := s.repository.GetByUserEmail(user.Oldemail)

	if err != nil {
		return err
	}

	err = dispatcher.SendChangeEmail(user)

	if err != nil {
		return err
	}

	err = s.redisRepository.SendCodeAuth(user.NewEmail, user.Code)

	if err != nil {
		return err
	}

	return nil
}

func (s *UserService) ChangePassword(user dto.UserUpdatePassword) error {

	u, err := s.repository.GetByUserEmail(user.Email)

	if err != nil {
		return err
	}

	svc := Encrypt{}

	passwordMatch, err := svc.CompareEncryptPassword(user.OldPassword, u.Password)

	if err != nil {
		return err
	}

	if !passwordMatch {
		return errors.New("incorrect password or username")
	}

	pass, err := svc.EncryptPassword(user.NewPassword)

	if err != nil {
		return err
	}

	err = s.repository.ChangePassword(pass, user.Email)

	if err != nil {
		return err
	}

	return nil
}

func (s *UserService) LostPassword(email string) error {

	dispatcher := MailDispatcherService{}

	var user dto.UserUpdatePassword
	user.Email = email

	_, err := s.repository.GetByUserEmail(email)

	if err != nil {
		return err
	}

	svc := RandomGenerator{}

	user.NewPassword = svc.GenerateRandomString(10)

	service := Encrypt{}

	pass, err := service.EncryptPassword(user.NewPassword)

	if err != nil {
		return err
	}

	err = s.repository.ChangePassword(pass, user.Email)

	if err != nil {
		return err
	}

	err = dispatcher.SendLostPasswordEmail(user)

	if err != nil {
		return err
	}

	return nil

}

func (s *UserService) ConfirmEmailChange(newEmail, oldEmail, code string) error {

	_, err := s.redisRepository.CheckCode(newEmail, code)

	if err != nil {
		return err
	}

	err = s.repository.ChangeEmail(newEmail, oldEmail)

	if err != nil {
		return err
	}

	return nil
}
