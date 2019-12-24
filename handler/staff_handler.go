package handler

import (
	"final_project/db"
	"github.com/labstack/echo/v4"
	"net/http"
)

func AddStudent(c echo.Context) error {
	var student db.Student
	if err := c.Bind(&student); err != nil {
		return c.JSON(http.StatusBadRequest, db.Error{Code: http.StatusBadRequest, Msg: "bad request"})
	}
	res, err := db.AddStudent(&student)
	if err != nil {
		return c.JSON(http.StatusBadRequest, db.Error{Code: http.StatusBadRequest, Msg: "bad request"})
	}
	return c.JSON(http.StatusOK, res)
}

func UpdateStudent(c echo.Context) error {
	var student db.Student
	if err := c.Bind(&student); err != nil {
		return c.JSON(http.StatusBadRequest, db.Error{Code: http.StatusBadRequest, Msg: "bad request"})
	}
	res, err := db.UpdateStudent(&student)
	if err != nil {
		return c.JSON(http.StatusBadRequest, db.Error{Code: http.StatusBadRequest, Msg: "bad request"})
	}
	return c.JSON(http.StatusOK, res)
}
