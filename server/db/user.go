package db

import (
	"server/log"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	OpenId      string `gorm:"column:openId"`
	Integral    int    `gorm:"column:integral"`
	IsVip       bool   `gorm:"column:is_vip"`
	AccessToken string `gorm:"column:access_token"`
	Status      string `gorm:"column:status"`
}

func (User) TableName() string {
	return "user"
}

func (user User) UserIsExistByOpenId() bool {
	var count int64

	if err := MYSQL.Model(user).Where("openid", user.OpenId).Count(&count).Error; err != nil {
		log.Log.Errorln(err)
		return false
	}

	if count > 0 {
		return true
	} else {
		return false
	}

}

func (user *User) FindUserById(id string) error {

	if err := MYSQL.First(&user, id).Error; err != nil {
		return err
	}

	return nil
}

func (user *User) FindUserByOpenId() error {
	if err := MYSQL.First(user, "openid = ?", user.OpenId).Error; err != nil {
		return err
	}

	return nil
}

func (user User) AddUser() error {
	if err := MYSQL.Create(user).Error; err != nil {
		log.Log.Errorf("user add err:%e", err)
		return err
	}

	return nil
}

func (user User) UpdateUser() error {

	if err := MYSQL.Save(user).Error; err != nil {
		return err
	}

	return nil
}

func (user User) DeleteUserByOpenId() error {

	if err := MYSQL.Where("openId = ?", user.OpenId).Delete(user).Error; err != nil {
		return err
	}

	return nil
}
