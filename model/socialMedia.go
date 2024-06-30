package model

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type SocialMedia struct {
	GormModel
	Name           string `gorm:"not null" json:"name"`
	SocialMediaUrl string `gorm:"type:text;not null" json:"social_media_url"`
	UserID         uint
	User           *User
}

func (u *SocialMedia) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}

func (u *SocialMedia) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errUpdate := govalidator.ValidateStruct(u)

	if errUpdate != nil {
		err = errUpdate
		return
	}

	err = nil
	return
}
