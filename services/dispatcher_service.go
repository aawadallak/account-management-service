package services

import (
	"encoding/json"
	"latest/app/kafka"
	"latest/config"
	"latest/dto"
)

type MailDispatcherService struct {
}

func (r *MailDispatcherService) SendEmail(user dto.User) error {

	service := kafka.NewKafkaProducer(config.GetConfig().KafkaTopicDispatcher)

	defer service.Close()

	x := dto.KafkaProducerRedis{
		Email: user.Email,
		Code:  user.Code,
	}

	t, err := json.Marshal(x)

	if err != nil {
		return err
	}

	err = service.Producer(t, "code")

	if err != nil {
		return err
	}

	return nil
}

func (r *MailDispatcherService) ReSendEmail(email, code string) error {

	service := kafka.NewKafkaProducer(config.GetConfig().KafkaTopicDispatcher)

	defer service.Close()

	x := dto.KafkaProducerRedis{
		Email: email,
		Code:  code,
	}

	t, err := json.Marshal(x)

	if err != nil {
		return err
	}

	err = service.Producer(t, "code")

	if err != nil {
		return err
	}

	return nil
}

func (r *MailDispatcherService) SendChangeEmail(user dto.UserUpdateEmail) error {

	service := kafka.NewKafkaProducer(config.GetConfig().KafkaTopicDispatcher)

	defer service.Close()

	x := dto.KafkaProducerRedis{
		Email: user.NewEmail,
		Code:  user.Code,
	}

	t, err := json.Marshal(x)

	if err != nil {
		return err
	}

	err = service.Producer(t, "change_email")

	if err != nil {
		return err
	}

	return nil
}

func (r *MailDispatcherService) SendLostPasswordEmail(user dto.UserUpdatePassword) error {

	service := kafka.NewKafkaProducer(config.GetConfig().KafkaTopicDispatcher)

	defer service.Close()

	x := dto.KafkaProducerLostPassword{
		Email:    user.Email,
		Password: user.NewPassword,
	}

	t, err := json.Marshal(x)

	if err != nil {
		return err
	}

	err = service.Producer(t, "lost_password")

	if err != nil {
		return err
	}

	return nil
}
