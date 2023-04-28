package repository

import "final-project/model"

type SocialMediaRepo interface {
	GetAllSocialMedia(in model.SocialMedia) (res []model.SocialMedia, err error)
	GetOneSocialMedia(id int64) (res model.SocialMedia, err error)
	CreateSocialMedia(in model.SocialMedia) (res model.SocialMedia, err error)
	UpdateSocialMedia(in model.SocialMedia) (res model.SocialMedia, err error)
	DeleteSocialMedia(id int64) (err error)
}

func (r Repo) GetAllSocialMedia(in model.SocialMedia) (res []model.SocialMedia, err error) {

	err = r.gorm.Model(&res).Where("user_id = ?", &in.UserID).Error

	if err != nil {
		return res, err
	}

	return res, nil
}

func (r Repo) GetOneSocialMedia(id int64) (res model.SocialMedia, err error) {

	err = r.gorm.First(&res, id).Error
	if err != nil {
		return res, err
	}

	return res, nil
}

func (r Repo) CreateSocialMedia(in model.SocialMedia) (res model.SocialMedia, err error) {

	err = r.gorm.Create(&in).Scan(&res).Error
	if err != nil {
		return res, err
	}

	return res, nil
}

func (r Repo) UpdateSocialMedia(in model.SocialMedia) (res model.SocialMedia, err error) {
	err = r.gorm.Model(&res).Where("id = ?", in.ID).Updates(model.SocialMedia{
		Name:           in.Name,
		SocialMediaUrl: in.SocialMediaUrl,
	}).Scan(&res).Error

	if err != nil {
		return res, nil
	}

	return res, nil

}

func (r Repo) DeleteSocialMedia(id int64) (err error) {

	in := model.SocialMedia{}

	err = r.gorm.Model(&in).Where("id = ?", id).Delete(&in).Error
	if err != nil {
		return err
	}

	return nil
}
