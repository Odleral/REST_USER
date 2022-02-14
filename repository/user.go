package repository

import (
	"gitlab.com/odleral/geoportal-go/model"
	"gorm.io/gorm"
)

func ListUser(db *gorm.DB) (model.Users, error) {
	users := make([]*model.User, 0)
	if err := db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func FindUser(db *gorm.DB, id string) (*model.User, error) {
	user := &model.User{}
	if err := db.Where("uuid = ?", id).Find(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func DeleteUser(db *gorm.DB, id string) error {
	user := &model.User{}

	if err := db.Where("uuid = ?", id).Delete(&user).Error; err != nil {
		return err
	}

	return nil
}

func CreateUser(db *gorm.DB, user *model.User) (*model.User, error) {
	if err := db.Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func UpdateUser(db *gorm.DB, user *model.User) error{
	if err := db.First(&model.User{}, user.UUID).Updates(user).Error; err != nil {
		return err
	}

	return nil
}
