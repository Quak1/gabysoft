package main

import (
	"context"
	"database/sql"
	_ "embed"
	"encoding/csv"
	"fmt"
	// "gabysoft/internal/config"
	"gabysoft/internal/database"
	"gabysoft/internal/queries"
	"log"
	"os"
	"path/filepath"
	"strconv"

	_ "modernc.org/sqlite"
)

func main() {
	db, err := initDB()
	if err != nil {
		log.Println(err)
		return
	}
	defer db.Close()

	data, err := readCSV("data.csv")
	if err != nil {
		log.Println(err)
		return
	}

	q := database.New(db)

	count, err := q.CountProducts(context.Background())
	if err != nil {
		log.Println(err)
		return
	}

	if count > 0 {
		fmt.Printf("Database has %d products\n", count)
		return
	}

	for _, record := range data[1:] {
		if err = processRecord(record, q); err != nil {
			log.Println(err)
			return
		}
		count++
	}

	fmt.Printf("Created %d database product entries\n", count)
}

func initDB() (*sql.DB, error) {
	// configDir, err := os.UserConfigDir()
	// if err != nil {
	// 	return nil, fmt.Errorf("Could not get user config dir: %w", err)
	// }
	//
	// appDir := filepath.Join(configDir, config.AppName)
	// if err := os.MkdirAll(appDir, 0755); err != nil {
	// 	return nil, fmt.Errorf("Could not create app directory: %w", err)
	// }

	appDir, _ := os.Getwd()

	dbPath := filepath.Join(appDir, "data.db")

	fmt.Println(dbPath)

	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, fmt.Errorf("Failed to open database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("Failed to ping database: %w", err)
	}

	if _, err := db.Exec(queries.Schema); err != nil {
		return nil, fmt.Errorf("Failed to apply schema: %w", err)
	}

	return db, nil
}

func readCSV(filename string) ([][]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		return nil, err
	}

	return records, nil
}

func processRecord(row []string, q *database.Queries) error {
	price, err := strconv.ParseFloat(row[4], 64)
	if err != nil {
		return err
	}

	err = q.CreateProduct(context.Background(), database.CreateProductParams{
		Name:        row[0],
		Code:        handleEmpty(row[1]),
		Barcode:     handleEmpty(row[2]),
		Description: handleEmpty(row[3]),
		Price:       price,
	})

	if err != nil {
		return err
	}

	return nil
}

func handleEmpty(entry string) sql.NullString {
	out := sql.NullString{
		String: entry,
		Valid:  true,
	}

	if entry == "" {
		out.Valid = false
	}

	return out
}
