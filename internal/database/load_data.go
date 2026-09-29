package database

import (
	"context"
	"database/sql"
	"encoding/csv"
	"fmt"
	"gabysoft/internal/store"
	"log"
	"os"
	"strconv"
)

func populateDB(db *sql.DB) {
	data, err := readCSV("data.csv")
	if err != nil {
		log.Println(err)
		return
	}

	q := store.New(db)

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

func processRecord(row []string, q *store.Queries) error {
	price, err := strconv.ParseFloat(row[4], 64)
	if err != nil {
		return err
	}

	err = q.CreateProduct(context.Background(), store.CreateProductParams{
		Name:        row[0],
		Code:        row[1],
		Barcode:     row[2],
		Description: row[3],
		Price:       price,
	})

	if err != nil {
		return err
	}

	return nil
}
