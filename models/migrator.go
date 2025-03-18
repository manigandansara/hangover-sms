package models

import (
	"github.com/robertantonyjaikumar/hangover-common/database"
	"github.com/robertantonyjaikumar/hangover-common/logger"
	"go.uber.org/zap"
	models "hangover/models/utils"
)

func GetTables() []interface{} {
	return []interface{}{
		&User{},
		&Role{},
		&Permission{},
		&UserRole{},
		&RolePermission{},
		&UserGroup{},
	}
}

func MigrateDB() {
	var tables []interface{}
	tables = append(tables, GetTables()...)
	migrations := database.Migrations{
		DB:     database.Db,
		Models: tables,
	}
	database.RunMigrations(migrations)
}

func SeedDB() {
	seed := []models.Seed{
		{Model: &[]UserGroup{}, FileName: "user_groups.json", CreateFunc: SeedUserGroup},
		{Model: &[]Role{}, FileName: "roles.json", CreateFunc: SeedRole},
		{Model: &[]User{}, FileName: "users.json", CreateFunc: SeedUser},
	}
	for _, s := range seed {
		if err := models.SeedModel(s.FileName, s.Model, s.CreateFunc); err != nil {
			logger.Fatal("Error seeding roles", zap.Error(err))
		}
	}
}
