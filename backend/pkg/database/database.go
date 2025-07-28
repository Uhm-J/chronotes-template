package database

import (
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DB holds the database connection
type DB struct {
	*gorm.DB
}

// New creates a new PostgreSQL database connection with GORM
func New(dsn string, level logger.LogLevel) (*DB, error) {
	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		PrepareStmt:            true,
		SkipDefaultTransaction: false,
		Logger:                 logger.Default.LogMode(level),
	})
	if err != nil {
		return nil, err
	}

	sqlDB, err := gormDB.DB()
	if err != nil {
		return nil, err
	}

	// Set connection pool settings
	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return &DB{gormDB}, nil
}

// Close closes the database connection
func (db *DB) Close() error {
	sqlDB, err := db.DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

// Migrate runs database migrations using GORM AutoMigrate
func (db *DB) Migrate(models ...interface{}) error {
	return db.AutoMigrate(models...)
}

// BeginTx starts a new transaction
func (db *DB) BeginTx() (*gorm.DB, error) {
	return db.Begin(), nil
}
