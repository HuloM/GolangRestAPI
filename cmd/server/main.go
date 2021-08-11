package main

import (
	"fmt"
	"github.com/HuloM/GolangRestAPI/internal/comment"
	"github.com/HuloM/GolangRestAPI/internal/database"
	"net/http"

	transportHTTP "github.com/HuloM/GolangRestAPI/internal/transport/http"
)

// App - the struct that contains things like pointers
// to database connections
type App struct {

}

// Run - sets up our application
func (app *App) Run() error {
	fmt.Println("Setting up our App")

	var err error
	db, err := database.NewDatabase()
	if err != nil {
		return err
	}

	err = database.MigrateDB(db)
	if err != nil {
		return err
	}

	commentService := comment.NewService(db)

	handler := transportHTTP.NewHandler(commentService)
	handler.SetupRoutes()

	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		fmt.Println("failed to set up server")
		return err
	}

	return nil
}

func main() {
	fmt.Println("go REST API")

	app := App{}

	if err := app.Run(); err != nil{
		fmt.Println("error starting up our REST API")
		fmt.Println(err)
	}
}
