package db

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"geth/internal/config"
	"geth/internal/model"
)

func Open(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)
	g, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn),
	})
	if err != nil {
		return nil, err
	}
	sqlDB, err := g.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(50)
	sqlDB.SetConnMaxLifetime(2 * time.Hour)
	sqlDB.SetConnMaxIdleTime(30 * time.Minute)

	if err := g.AutoMigrate(
		&model.ChainCursor{},
		&model.UUPSAuction{},
		&model.UUPSBid{},
	); err != nil {
		return nil, fmt.Errorf("auto migrate: %w", err)
	}
	return g, nil
}
