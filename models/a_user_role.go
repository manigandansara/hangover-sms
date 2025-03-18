package models

type UserRole struct {
	PreModel
	UserID   uint  `gorm:"primaryKey"`
	RoleID   uint  `gorm:"primaryKey"`
	IsActive *bool `json:"is_active"`
}

func (UserRole) TableName() string {
	return "user_roles"
}
