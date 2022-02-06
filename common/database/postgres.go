package database

import (
	"github.com/yusuftatli/go-template-with-redis-gorm-api/entities"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB(databaseUrl string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(databaseUrl), &gorm.Config{})
	if err != nil {
		zap.S().Fatal(
			"Database connection failed",
			zap.Error(err),
		)
		return nil
	}
	return db
}

func CloseDBConnection(db *gorm.DB) {
	con, err := db.DB()
	if err != nil {
		zap.S().Fatal("Database connection could not get", zap.Error(err))
		return
	}
	if err := con.Close(); err != nil {
		zap.S().Fatal("Database connection could not closed", zap.Error(err))
	}
}

func AutoMigrate(db *gorm.DB) {
	_ = db.AutoMigrate(entities.Template{})
	_ = db.AutoMigrate(entities.User{})
}
