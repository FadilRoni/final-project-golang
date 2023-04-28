package service

import "final-project/repository"

type Service struct {
	repo repository.RepoInterface
}

type ServiceInterface interface {
	UserService
	SocialMediaService
	PhotoService
	CommentService
}

func NewService(repo repository.RepoInterface) ServiceInterface {
	return &Service{repo: repo}
}
