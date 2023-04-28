package service

import "final-project/model"

type SocialMediaService interface {
	GetAllSocialMedia(in model.SocialMedia) (res []model.SocialMedia, err error)
	GetOneSocialMedia(id int64) (res model.SocialMedia, err error)
	CreateSocialMedia(in model.SocialMedia) (res model.SocialMedia, err error)
	UpdateSocialMedia(in model.SocialMedia) (res model.SocialMedia, err error)
	DeleteSocialMedia(id int64) (err error)
}

func (s *Service) GetAllSocialMedia(in model.SocialMedia) (res []model.SocialMedia, err error) {
	return s.repo.GetAllSocialMedia(in)
}

func (s *Service) GetOneSocialMedia(id int64) (res model.SocialMedia, err error) {
	return s.repo.GetOneSocialMedia(id)
}

func (s *Service) CreateSocialMedia(in model.SocialMedia) (res model.SocialMedia, err error) {
	res, err = s.repo.CreateSocialMedia(in)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (s *Service) UpdateSocialMedia(in model.SocialMedia) (res model.SocialMedia, err error) {
	return s.repo.UpdateSocialMedia(in)
}

func (s *Service) DeleteSocialMedia(id int64) (err error) {
	return s.repo.DeleteSocialMedia(id)
}
