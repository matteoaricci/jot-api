package journals

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/matteoaricci/jot-api/models/journal"
	"github.com/matteoaricci/jot-api/service/journal"
	echoSwagger "github.com/swaggo/echo-swagger"
	"net/http"
)

// @title Journal API
// @version 1.0
// @description This is a Journal management API
// @host localhost:8080
// @BasePath /api

func AddRoutes(e *echo.Echo) {
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.GET("/api/journals", getAllJournals)
	e.POST("/api/journals", createJournal)
	e.GET("/api/journals/:id", getJournalByID)
	e.DELETE("/api/journals/:id", deleteJournal)
	e.PUT("/api/journals/:id", updateJournal)
}

// @Summary Get all journals
// @Description Get a list of all journals
// @Tags journals
// @Accept json
// @Produce json
// @Success 200 {array} models.JournalVM[]
// @Failure 500 {object} echo.HTTPError
// @Router /journals [get]
func getAllJournals(c echo.Context) error {
	j, err := journal.All()
	if err != nil {
		return c.JSON(err.Code, err)
	}
	return c.JSON(http.StatusOK, j)
}

// @Summary Get a journal
// @Description Get a journal by ID
// @Tags journals
// @Accept json
// @Produce json
// @Param id path string true "Journal ID"
// @Success 200 {object} models.JournalVM
// @Failure 404 {object} echo.HTTPError
// @Failure 500 {object} echo.HTTPError
// @Router /journals/{id} [get]
func getJournalByID(c echo.Context) error {
	id := c.Param("id")
	j, err := journal.Get(id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, *j)
}

// @Summary Create a journal
// @Description Create a new journal entry
// @Tags journals
// @Accept json
// @Produce json
// @Param journal body models.CreateOrPutJournalVM true "Journal object"
// @Success 201 {string} string "Journal ID"
// @Failure 400 {object} echo.HTTPError
// @Failure 500 {object} echo.HTTPError
// @Router /journals [post]
func createJournal(c echo.Context) error {
	e := echo.New()
	e.Validator = &models.CreateOrPutJournalValidator{Validator: validator.New()}
	var j models.CreateOrPutJournalVM
	err := c.Bind(&j)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	if err = c.Validate(&j); err != nil {
		return err
	}
	newJID, httpErr := journal.Create(j)
	if httpErr != nil {
		return httpErr
	}
	return c.JSON(http.StatusCreated, *newJID)
}

// @Summary Update a journal
// @Description Update a journal by ID
// @Tags journals
// @Accept json
// @Produce json
// @Param id path string true "Journal ID"
// @Param journal body models.CreateOrPutJournalVM true "Journal object"
// @Success 200 {object} models.JournalVM
// @Failure 400 {object} echo.HTTPError
// @Failure 404 {object} echo.HTTPError
// @Failure 500 {object} echo.HTTPError
// @Router /journals/{id} [put]
func updateJournal(c echo.Context) error {
	e := echo.New()
	e.Validator = &models.CreateOrPutJournalValidator{Validator: validator.New()}
	id := c.Param("id")
	var j models.CreateOrPutJournalVM
	err := c.Bind(&j)
	if err != nil {
		return err
	}
	if err = c.Validate(&j); err != nil {
		return err
	}
	newJ, httpErr := journal.Put(id, j)
	if httpErr != nil {
		return httpErr
	}
	return c.JSON(http.StatusOK, newJ)
}

// @Summary Delete a journal
// @Description Delete a journal by ID
// @Tags journals
// @Accept json
// @Produce json
// @Param id path string true "Journal ID"
// @Success 204 "No Content"
// @Failure 404 {object} echo.HTTPError
// @Failure 500 {object} echo.HTTPError
// @Router /journals/{id} [delete]
func deleteJournal(c echo.Context) error {
	id := c.Param("id")
	err := journal.Delete(id)
	if err != nil {
		return err
	}
	return c.NoContent(http.StatusNoContent)
}
