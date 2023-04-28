package repository

import "final-project/model"

type CommentRepo interface {
	GetAllComment(in model.Comment) (res []model.Comment, err error)
	GetOneComment(in model.Comment) (res model.Comment, err error)
	CreateComment(in model.Comment) (res model.Comment, err error)
	UpdateComment(in model.Comment) (res model.Comment, err error)
	DeleteComment(in model.Comment) (err error)
}

func (r Repo) GetAllComment(in model.Comment) (res []model.Comment, err error) {

	err = r.gorm.Model(&res).Where("user_id = ?", &in.UserID).Error

	if err != nil {
		return res, err
	}

	return res, nil
}

func (r Repo) GetOneComment(in model.Comment) (res model.Comment, err error) {

	err = r.gorm.First(&res, in.ID).Error
	if err != nil {
		return res, err
	}

	return res, nil
}

func (r Repo) CreateComment(in model.Comment) (res model.Comment, err error) {

	err = r.gorm.Create(&in).Scan(&res).Error
	if err != nil {
		return res, err
	}

	return res, nil
}

func (r Repo) UpdateComment(in model.Comment) (res model.Comment, err error) {
	err = r.gorm.Model(&res).Where("id = ?", in.ID).Updates(model.Comment{
		Message: in.Message,
	}).Error

	if err != nil {
		return res, nil
	}

	return res, nil

}

func (r Repo) DeleteComment(in model.Comment) (err error) {

	err = r.gorm.Model(&in).Where("id = ?", in.ID).Delete(&in).Error
	if err != nil {
		return err
	}

	return nil
}
