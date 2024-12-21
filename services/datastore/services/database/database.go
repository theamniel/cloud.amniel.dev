package database

import (
	"cloud.amniel.dev/services/datastore/services/bans"
	"github.com/knadh/koanf/v2"
	"go.uber.org/zap"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func New(k *koanf.Koanf, log *zap.SugaredLogger) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(k.String("database.path")), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		Logger:                 logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	if err := db.AutoMigrate(&bans.Ban{}); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	log.Info("database connection established")
	return db
}

func IsNotFound(err error) bool {
	return err == gorm.ErrRecordNotFound
}
