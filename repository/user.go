package repository

import "final-project/model"

type UserRepo interface {
	Register(in model.User) (res model.User, err error)
	Login(in model.User) (res model.User, err error)
}

func (r Repo) Register(in model.User) (res model.User, err error) {

	err = r.gorm.Create(&in).Scan(&res).Error
	if err != nil {
		return res, err
	}

	return res, nil
}

func (r Repo) Login(in model.User) (res model.User, err error) {

	err = r.gorm.Where("email = ?", in.Email).Take(&res).Error
	if err != nil {
		return res, err
	}

	return res, nil
}
