package service

import "final-project/model"

type CommentService interface {
	GetAllComment(in model.Comment) (res []model.Comment, err error)
	GetOneComment(in model.Comment) (res model.Comment, err error)
	CreateComment(in model.Comment) (res model.Comment, err error)
	UpdateComment(in model.Comment) (res model.Comment, err error)
	DeleteComment(in model.Comment) (err error)
}

func (s *Service) GetAllComment(in model.Comment) (res []model.Comment, err error) {
	return s.repo.GetAllComment(in)
}

func (s *Service) GetOneComment(in model.Comment) (res model.Comment, err error) {
	return s.repo.GetOneComment(in)
}

func (s *Service) CreateComment(in model.Comment) (res model.Comment, err error) {
	res, err = s.repo.CreateComment(in)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (s *Service) UpdateComment(in model.Comment) (res model.Comment, err error) {
	return s.repo.UpdateComment(in)
}

func (s *Service) DeleteComment(in model.Comment) (err error) {
	return s.repo.DeleteComment(in)
}
