package storage

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB represents a DB connection that can be used to run SQL queries.
type DB struct {
	db *gorm.DB
}

// NewDatabase returns a new DB connection that wraps the given gorm.DB instance.
func NewDatabase(db *gorm.DB) *DB {
	return &DB{db}
}

// DB returns the gorm.DB wrapped by this object.
func (db *DB) DB() *gorm.DB {
	return db.db
}

// With returns a Builder that can be used to build and execute SQL queries.
// With will return the transaction if it is found in the given context.
// Otherwise, it will return a DB connection associated with the context.
func (db *DB) With(c *gin.Context) *gorm.DB {
	return db.db.WithContext(c)
}

func LoadAllTables(db *gorm.DB) error {
	return nil
}

func ConnectToDatabase() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open("postgres://postgres:postgres@localhost:5432/hub_dev"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = LoadAllTables(db)
	if err != nil {
		return nil, err
	}
	return db, nil
}
