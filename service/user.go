package service

import (
	"errors"
	"final-project/model"
)

type UserService interface {
	Register(in model.User) (res model.User, err error)
	Login(in model.User) (res model.User, err error)
}

func (s *Service) Register(in model.User) (res model.User, err error) {

	if int(in.Age) < 8 {
		return res, errors.New("field age value must more than 8")
	}

	res, err = s.repo.Register(in)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (s *Service) Login(in model.User) (res model.User, err error) {
	res, err = s.repo.Login(in)
	if err != nil {
		return res, err
	}

	return res, nil
}
