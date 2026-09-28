package main

import (
	"context"
	"database/sql"
	"fmt"
	"gabysoft/internal/database"
	"gabysoft/internal/store"
	"log"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx   context.Context
	db    *sql.DB
	query *store.Queries
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx

	if err := database.DeleteDB(); err != nil {
		log.Println(err)
		runtime.Quit(ctx)
	}

	db, err := database.InitDB()
	if err != nil {
		log.Println(err)
		runtime.Quit(ctx)
	}
	a.db = db

	q := store.New(db)
	a.query = q
}

func (a *App) shutdown(ctx context.Context) {
	a.db.Close()
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) GetAllProducts() ([]store.Product, error) {
	return a.query.GetProducts(context.Background())
}
