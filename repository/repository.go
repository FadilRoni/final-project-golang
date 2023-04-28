package repository

import "gorm.io/gorm"

type Repo struct {
	gorm *gorm.DB
}

type RepoInterface interface {
	UserRepo
	SocialMediaRepo
	PhotoRepo
	CommentRepo
}

func NewRepo(gorm *gorm.DB) *Repo {
	return &Repo{gorm: gorm}
}
