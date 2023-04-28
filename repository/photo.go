package repository

import (
	"final-project/model"
)

type PhotoRepo interface {
	GetAllPhoto(in model.Photo) (res []model.Photo, err error)
	GetOnePhoto(id int64) (res model.Photo, err error)
	CreatePhoto(in model.Photo) (res model.Photo, err error)
	UpdatePhoto(in model.Photo) (res model.Photo, err error)
	DeletePhoto(id int64) (err error)
}

func (r Repo) GetAllPhoto(in model.Photo) (res []model.Photo, err error) {

	err = r.gorm.Model(&res).Where("user_id = ?", &in.UserID).Error

	if err != nil {
		return res, err
	}

	return res, nil
}

func (r Repo) GetOnePhoto(id int64) (res model.Photo, err error) {

	// c := model.Comment{}

	err = r.gorm.Preload("Comments", id).First(&res, id).Error
	if err != nil {
		return res, err
	}

	return res, nil
}

func (r Repo) CreatePhoto(in model.Photo) (res model.Photo, err error) {

	err = r.gorm.Create(&in).Scan(&res).Error
	if err != nil {
		return res, err
	}

	return res, nil
}

func (r Repo) UpdatePhoto(in model.Photo) (res model.Photo, err error) {
	err = r.gorm.Model(&res).Where("id = ?", in.ID).Updates(model.Photo{
		Title:    in.Title,
		Caption:  in.Caption,
		PhotoUrl: in.PhotoUrl,
	}).Error

	if err != nil {
		return res, nil
	}

	return res, nil

}

func (r Repo) DeletePhoto(id int64) (err error) {

	in := model.Photo{}

	err = r.gorm.Model(&in).Where("id = ?", id).Delete(&in).Error
	if err != nil {
		return err
	}

	return nil
}
