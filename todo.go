package main

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/matteo107/go-echo-vue/handlers"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Db
	db := initDB("storage.db")
	migrate(db)

	// Echo instance
	e := echo.New()
	e.HideBanner = true
	e.Logger.SetLevel(99)

	// Middleware
	e.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
	}))
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", func(c echo.Context) error { return c.JSON(200, "Hello") })

	e.GET("/tasks", handlers.GetTasks(db))
	e.PUT("/tasks", handlers.PutTask(db))
	e.DELETE("/tasks/:id", handlers.DeleteTask(db))
	// Start server
	e.Logger.Fatal(e.Start(":8000"))

}

func initDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)

	// Here we check for any db errors then exit
	if err != nil {
		panic(err)
	}

	// If we don't get any errors but somehow still don't get a db connection
	// we exit as well
	if db == nil {
		panic("db nil")
	}
	return db
}

func migrate(db *sql.DB) {
	sql := `
    CREATE TABLE IF NOT EXISTS tasks(
        id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        name VARCHAR NOT NULL
    );
    `

	_, err := db.Exec(sql)
	// Exit if something goes wrong with our SQL statement above
	if err != nil {
		panic(err)
	}
}
