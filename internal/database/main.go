package database

import (
	"database/sql"
	"fmt"
	"gabysoft/internal/config"
	"os"
	"path/filepath"

	_ "modernc.org/sqlite"
)

func InitDB() (*sql.DB, error) {
	dbPath, err := getDBPath()
	if err != nil {
		return nil, err
	}

	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, fmt.Errorf("Failed to open database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("Failed to ping database: %w", err)
	}

	if _, err := db.Exec(schema); err != nil {
		return nil, fmt.Errorf("Failed to apply schema: %w", err)
	}

	populateDB(db)

	return db, nil
}

func DeleteDB() error {
	dbPath, err := getDBPath()
	if err != nil {
		return err
	}

	if err = os.RemoveAll(dbPath); err != nil {
		return fmt.Errorf("Failed to delete db file")
	}

	return nil
}

func getDBPath() (string, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return "", fmt.Errorf("Could not get user config dir: %w", err)
	}

	appDir := filepath.Join(configDir, config.AppName)
	if err := os.MkdirAll(appDir, 0755); err != nil {
		return "", fmt.Errorf("Could not create app directory: %w", err)
	}

	dbPath := filepath.Join(appDir, "data.db")

	fmt.Println(dbPath)

	return dbPath, nil
}
