package db

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/storygraph/story-graph/pkg/config"
	"github.com/storygraph/story-graph/pkg/db/dao"
)

var db *gorm.DB

func init() {
	if db != nil {
		return
	}

	if err := connectToDB(); err != nil {
		log.Fatalf("Error connecting to DB: %s", err.Error())
	}

	if err := migrateDB(); err != nil {
		log.Fatalf("Error during DB migration: %s", err.Error())
	}
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

	log.Printf("Successfully connected to DB %s on %s:%d...", cfg.DBName, cfg.DBHost, cfg.DBPort)
	return nil
}

func migrateDB() error {
	// Using null values as of only the data model structures are needed.
	dataModels := []dao.DataModel{
		dao.NewStory(db, nil),
		dao.NewWeenie(db, nil, nil, nil, 0),
		dao.NewRelationType(db, nil),
		dao.NewRelationDelta(db, nil, false, 0, 0, 0),
		dao.NewWeenieDelta(db, nil, nil, 0, 0, false, 0),
		dao.NewPossessionDelta(db, nil, nil, false, 0),
		dao.NewTagDelta(db, nil, nil, false, 0),
		dao.NewEvent(db, nil, 0, 0),
		dao.NewAction(db, nil, []uint{}, 0, 0, 0),
		dao.NewSection(db, nil, 0, 0),
	}

	for _, dataModel := range dataModels {
		if err := dataModel.Migrate(); err != nil {
			return err
		}
	}

	log.Printf("Successfully migrated DB schema...")
	return nil
}
