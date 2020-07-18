package db

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/storygraph/story-graph/pkg/config"
	"github.com/storygraph/story-graph/pkg/db/dto"
)

var db *gorm.DB

func init() {
	if db != nil {
		return
	}

	if err := connectToDB(); err != nil {
		log.Fatalf("Error connecting to db: %s", err.Error())
	}

	migrateDB()
}

func GetConn() *gorm.DB {
	return db
}

func connectToDB() (err error) {
	cfg := config.GetConfig()
	dbConnStr := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s", cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBName, cfg.DBPass)

	db, err = gorm.Open("postgres", dbConnStr)
	if err != nil {
		return err
	}

	return nil
}

func migrateDB() (err error) {
	dataModels := []dto.DataModel{
		&dto.RelationType{},
		&dto.Relation{},
	}

	for _, dataModel := range dataModels {
		dataModel.Migrate(db)
	}

	return nil
}
