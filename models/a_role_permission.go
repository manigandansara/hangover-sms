package models

type RolePermission struct {
	PreModel
	RoleID       uint  `gorm:"primaryKey"`
	PermissionID uint  `gorm:"primaryKey"`
	IsActive     *bool `json:"is_active"`
}
