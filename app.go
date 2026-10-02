package main

import (
	"context"
	"database/sql"
	"gabysoft/internal/database"
	"gabysoft/internal/store"
	"gabysoft/internal/tasks"
	"log"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx   context.Context
	db    *sql.DB
	tasks *tasks.Tasks
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{
		tasks: tasks.NewTasks(),
	}
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
	a.tasks.Inject(ctx, q)
}

func (a *App) shutdown(ctx context.Context) {
	a.db.Close()
}

func (a *App) ShowMessage(title, message string) {
	_, err := runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
		Type:    runtime.InfoDialog,
		Title:   title,
		Message: message,
	})

	if err != nil {
		log.Println(err)
	}
}

func (a *App) ShowConfirm(title, message string) bool {
	result, err := runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
		Type:          runtime.QuestionDialog,
		Title:         title,
		Message:       message,
		Buttons:       []string{"Yes", "No"},
		DefaultButton: "No",
		CancelButton:  "No",
	})
	if err != nil {
		log.Println(err)
		return false
	}

	if result == "Yes" {
		return true
	}

	return false
}
