package model

import (
	"final-project/helper"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	GormModel
	Username    string        `gorm:"not null;uniqueIndex" json:"username" form:"username" valid:"required~Field Username is required"`
	Email       string        `gorm:"not null;uniqueIndex" form:"email" json:"email" valid:"required~Field Email is required,email~Invalid email format"`
	Password    string        `gorm:"not null" json:"password" form:"password" valid:"required~Field Password is required, minstringlength(6)~Password Length Minimum 6 characters"`
	Age         int           `gorm:"not null" json:"age" valid:"required~Field Age is required"`
	Photos      []Photo       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"photos"`
	Comments    []Comment     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"comments"`
	SocialMedia []SocialMedia `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"social_media"`
}

type Photo struct {
	GormModel
	Title    string `json:"title" gorm:"not null" form:"title" valid:"required~Field Title is required"`
	Caption  string `json:"caption" form:"caption"`
	PhotoUrl string `json:"photo_url" gorm:"not null" form:"photo_url" valid:"required~Field Caption is required,url~Invalid URL format"`
	UserID   int    `json:"user_id"`
	User     *User
	Comments []Comment `json:"comments" gorm:"foreignKey:PhotoID"`
}

type Comment struct {
	GormModel
	Message string `json:"message" gorm:"not null" form:"message" valid:"required~Field Message is required"`
	Photo   *Photo
	User    *User
	PhotoID int `json:"photo_id"`
	UserID  int `json:"user_id"`
}

type SocialMedia struct {
	GormModel
	Name           string `json:"name" gorm:"not null" form:"name" valid:"required~Field Name is required"`
	SocialMediaUrl string `json:"social_media_url" form:"social_media_url" gorm:"not null" valid:"required~Field Social Media URL is required,url~Invalid URL format"`
	UserID         int    `json:"user_id"`
	User           *User
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)
	if errCreate != nil {
		err = errCreate
		return
	}

	u.Password = helper.HassPass(u.Password)
	err = nil
	return
}

func (s *SocialMedia) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(s)
	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}

func (s *SocialMedia) BerforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(s)
	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}

func (p *Photo) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)
	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}

func (p *Photo) BerforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)
	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}

func (c *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(c)
	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}

func (c *Comment) BerforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(c)
	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}
