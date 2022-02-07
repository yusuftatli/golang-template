package database

import (
	"github.com/yusuftatli/go-template-with-redis-gorm-api/internal/entities"
	"github.com/yusuftatli/go-template-with-redis-gorm-api/pkg/logging"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var logger *logging.LogProvider

func GetPostgresDB() *gorm.DB {
	return db
}

func ConnectDB(databaseUrl string) *gorm.DB {
	logger = logging.GetLogger()
	db, err := gorm.Open(postgres.Open(databaseUrl), &gorm.Config{})
	if err != nil {
		logger.ZapLogger.Sugar().Fatal(
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
		logger.ZapLogger.Sugar().Fatal("Database connection could not get", zap.Error(err))
		return
	}
	if err := con.Close(); err != nil {
		logger.ZapLogger.Sugar().Fatal("Database connection could not closed", zap.Error(err))
	}
}

func AutoMigrate(db *gorm.DB) {
	_ = db.AutoMigrate(entities.Example{})
	_ = db.AutoMigrate(entities.User{})
}
