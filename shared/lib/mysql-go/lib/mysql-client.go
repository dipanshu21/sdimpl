package lib

import (
	"database/sql"
	"fmt"

	"encoding/json"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type DBConfig struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
}

func getDBConfig() (*DBConfig, error) {
	file, err := os.Open(DefaultDBConfigFile)
	if err != nil {
		return nil, fmt.Errorf("failed to open config file: %w", err)
	}
	defer file.Close()

	var config DBConfig
	if err := json.NewDecoder(file).Decode(&config); err != nil {
		return nil, fmt.Errorf("failed to decode config: %w", err)
	}

	return &config, nil
}

func getDBDSN() string {
	// Read config from db.config.json
	config, err := getDBConfig()
	if err != nil {
		return ""
	}

	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		config.Username, config.Password, config.Host, config.Port, DefaultDBName)
}

func createConnectionToDBInternal(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MySQL: %w", err)
	}

	// Check the connection
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping MySQL: %w", err)
	}

	return db, nil
}

const (
	// DefaultDBConfigFile is the default path to the database configuration file.
	DefaultDBConfigFile = "db.config.json"
	// DefaultDBName is the default name of the database.
	DefaultDBName = "db_exp"
)

var DB_DSN string = getDBDSN()

func CreateNewConnectionToDB() (*sql.DB, error) {
	if DB_DSN == "" {
		return nil, fmt.Errorf("database DSN is empty, please check your configuration")
	}

	db, err := createConnectionToDBInternal(DB_DSN)
	if err != nil {
		return nil, fmt.Errorf("error creating connection to database: %w", err)
	}

	return db, nil
}
