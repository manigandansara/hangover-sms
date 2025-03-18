package models

type Permission struct {
	PreModel
	ID          uint   `gorm:"primaryKey"`
	Name        string `json:"name" gorm:"unique;not null"`
	Description string `json:"description"`
	Roles       []Role `gorm:"many2many:role_permissions;"`
	IsActive    *bool  `json:"is_active"`
}

func (u *Permission) TableName() string {
	return "permissions"
}
