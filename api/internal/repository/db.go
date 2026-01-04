package repository

import (
	"fmt"

	"github.com/naotch/minibo/api/internal/config"
	"github.com/naotch/minibo/api/internal/model"
	"github.com/naotch/minibo/api/pkg/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	DB = newSQLFactory()
	err := DB.AutoMigrate(&model.User{})
	if err != nil {
		logger.Panic("Database migration failed", err)
	}
}

func newSQLFactory() *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tokyo",
		config.Config.DBHost,
		config.Config.DBUser,
		config.Config.DBPass,
		config.Config.DBName,
		config.Config.DBPort,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Panic("Database connection failed.", err)
	}
	return db
}
