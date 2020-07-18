package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

var cfg *StorygraphConfig

func init() {
	var err error

	if cfg != nil {
		return
	}

	cfg, err = newStoryGraphConfig()
	if err != nil {
		log.Fatalf("Error reading config: %s", err.Error())
	}
}

type StorygraphConfig struct {
	DBHost string
	DBPort uint64
	DBUser string
	DBPass string
	DBName string

	Port uint64
}

func GetConfig() *StorygraphConfig {
	return cfg
}

func newStoryGraphConfig() (cfg *StorygraphConfig, err error) {
	cfg = &StorygraphConfig{}

	if err = cfg.readDBConfig(); err != nil {
		return nil, err
	}

	if err = cfg.readPort(); err != nil {
		return nil, err
	}

	return cfg, nil
}

func (sc *StorygraphConfig) readDBConfig() (err error) {
	sc.DBHost = os.Getenv("DB_HOST")

	if sc.DBPort, err = readNumFromEnv("DB_PORT"); err != nil {
		return err
	}

	sc.DBUser = os.Getenv("DB_USER")
	sc.DBPass = os.Getenv("DB_PASS")
	sc.DBName = os.Getenv("DB_NAME")

	return nil
}

func (sc *StorygraphConfig) readPort() (err error) {
	if sc.Port, err = readNumFromEnv("PORT"); err != nil {
		return err
	}

	return nil
}

func readNumFromEnv(envVar string) (num uint64, err error) {
	if num, err = strconv.ParseUint(os.Getenv(envVar), 10, 64); err != nil {
		return 0, fmt.Errorf("Error reading db port: %s", err.Error())
	}

	return num, nil
}
