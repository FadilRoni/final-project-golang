package service

import "final-project/model"

type PhotoService interface {
	GetAllPhoto(in model.Photo) (res []model.Photo, err error)
	GetOnePhoto(id int64) (res model.Photo, err error)
	CreatePhoto(in model.Photo) (res model.Photo, err error)
	UpdatePhoto(in model.Photo) (res model.Photo, err error)
	DeletePhoto(id int64) (err error)
}

func (s *Service) GetAllPhoto(in model.Photo) (res []model.Photo, err error) {
	return s.repo.GetAllPhoto(in)
}

func (s *Service) GetOnePhoto(id int64) (res model.Photo, err error) {
	return s.repo.GetOnePhoto(id)
}

func (s *Service) CreatePhoto(in model.Photo) (res model.Photo, err error) {
	res, err = s.repo.CreatePhoto(in)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (s *Service) UpdatePhoto(in model.Photo) (res model.Photo, err error) {
	return s.repo.UpdatePhoto(in)
}

func (s *Service) DeletePhoto(id int64) (err error) {
	return s.repo.DeletePhoto(id)
}
