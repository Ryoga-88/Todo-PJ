package repository

import (
	"fmt"
	"log"

	"github.com/Ryoga-88/Todo-PJ/backend/config"
	"github.com/Ryoga-88/Todo-PJ/backend/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Models = []interface{}{
	&entity.User{},
	&entity.Task{},
}

func NewDB(config *config.Config) (*gorm.DB, error) {
	url := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", config.POSTGRES_USER, config.POSTGRES_PW, config.POSTGRES_HOST, config.POSTGRES_PORT, config.POSTGRES_DB)
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("error connecting to database: %w, %s", err, url)
	}
	fmt.Println("Connected to database")
	return db, nil
}

func CloseDB(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalln(err)
	}
	if err := sqlDB.Close(); err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Closed database connection")
}

func Migrate(db *gorm.DB) error {
	db.AutoMigrate(Models...)
	defer fmt.Println("Migration has been completed")

	return nil
}
