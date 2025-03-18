package models

import (
	"fmt"
	"github.com/robertantonyjaikumar/hangover-common/database"
	"github.com/robertantonyjaikumar/hangover-common/logger"
	"go.uber.org/zap"
	models "hangover/models/utils"
)

type User struct {
	PreModelWithUUID
	Username     string    `json:"username" gorm:"unique;not null"`
	Email        string    `json:"email" gorm:"unique;not null"`
	PasswordHash string    `json:"password_hash"`
	UserGroup    uint      `json:"user_group"`
	Group        UserGroup `gorm:"foreignKey:UserGroup"`
	Roles        []Role    `gorm:"many2many:user_roles;"`
	IsActive     *bool     `json:"is_active"`
}

func (u *User) TableName() string {
	return "users"
}

func GetUserByUserName(username string) (*User, error) {
	var user User
	if err := database.Db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func ValidateUserByUserNameAndPassword(username, password string) (*User, error) {
	user, err := GetUserByUserName(username)
	if err != nil {
		return nil, err
	}
	if models.ValidatePassword(user.PasswordHash, password) {
		return user, nil
	}
	return nil, fmt.Errorf("invalid username or password")
}

func SeedUser(model interface{}) error {
	users, ok := model.(*[]User)
	if !ok {
		return fmt.Errorf("invalid model type")
	}
	for _, user := range *users {
		user.PasswordHash, _ = models.HashPassword(user.PasswordHash)
		if err := database.Db.FirstOrCreate(&user, "username = ?", user.Username).Error; err != nil {
			logger.Error("Error creating user seed: "+user.Username, zap.Error(err))
			return err
		}
	}
	return nil
}
