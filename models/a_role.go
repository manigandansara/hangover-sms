package models

import (
	"fmt"
	"github.com/robertantonyjaikumar/hangover-common/database"
	"github.com/robertantonyjaikumar/hangover-common/logger"
	"go.uber.org/zap"
)

type Role struct {
	PreModel
	Name        string `gorm:"unique;not null"`
	Description string
	Permissions []Permission `gorm:"many2many:role_permissions;"`
	Users       []User       `gorm:"many2many:user_roles;"`
	IsActive    *bool        `json:"is_active"`
}

func (u *Role) TableName() string {
	return "roles"
}

func SeedRole(model interface{}) error {
	roles, ok := model.(*[]Role)
	if !ok {
		return fmt.Errorf("invalid model type")
	}
	for _, role := range *roles {
		if err := database.Db.FirstOrCreate(&role, "name = ?", role.Name).Error; err != nil {
			logger.Error("Error creating user seed: "+role.Name, zap.Error(err))
			return err
		}
	}
	return nil
}
