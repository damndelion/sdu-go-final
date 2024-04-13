// Package postgres implements postgres connection.
package postgres

import (
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Postgres -.
type Postgres struct {
	*gorm.DB
}

// New -.
func New(url string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

// Close -.
func Close(db *gorm.DB) {
	sqlDB, _ := db.DB()

	err := sqlDB.Close()
	if err != nil {
		logrus.Printf("error closing database: %s", err.Error())
		return
	}
}
