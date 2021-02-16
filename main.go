package main

import (
	"github.com/gustavopergola/go-rest-study/app"
	"os"
)

func main() {
	db := app.StartDB()

	if len(os.Args) > 1 && os.Args[1] == "migration" {
		app.Migrate(db)
		return
	}

	app.StartServer(db)
}
