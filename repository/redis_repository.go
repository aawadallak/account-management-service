package repository

import (
	"errors"
	"latest/app/redis"
	"strings"
	"time"
)

type RedisRepository struct{}

func (t *RedisRepository) SendCodeAuth(key, value string) error {

	s := redis.RedisClient()

	err := s.Writer(key, value, time.Minute*5)

	if err != nil {
		return err
	}

	return nil
}

func (t *RedisRepository) CheckCode(key string, code string) (bool, error) {

	s := redis.RedisClient()

	w, err := s.Find(key)

	w = strings.ReplaceAll(w, `"`, "")

	if err != nil {
		return false, errors.New("code already expired or it's wrong")
	}

	if w != code {
		return false, errors.New("invalid code. Try again")
	}

	return true, nil
}
