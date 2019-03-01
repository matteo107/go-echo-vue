// handlers/tasks.go
package handlers

import (
	"database/sql"
	"github.com/labstack/gommon/log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/matteo107/go-echo-vue/models"
)

type H map[string]interface{}

// GetTasks endpoint
func GetTasks(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, models.GetTasks(db))
	}
}

// PutTasks endpoint
func PutTask(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Instantiate a new task
		var task models.Task
		// Map imcoming JSON body to the new Task
		err := c.Bind(&task)
		if err != nil {
			log.Printf("Failed processing PutTask request: %s\n", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		// Add a task using our new model
		id, err := models.PutTask(db, task.Name)

		// Return a JSON response if successful
		if err == nil {
			log.Printf("This is your task: %#v\n", task)
			return c.JSON(http.StatusCreated, H{
				"created": id,
			})
			// Handle any errors
		} else {
			return err
		}
	}
}

// DeleteTask endpoint
func DeleteTask(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		_, err := models.DeleteTask(db, id)
		// Return a JSON response on success
		if err == nil {
			return c.JSON(http.StatusOK, H{
				"deleted": id,
			})
			// Handle errors
		} else {
			return err
		}
	}
}
