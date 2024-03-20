package controllers

import (
	"CobaEcho/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func FetchAllUsers(c echo.Context) error {
	result, err := models.FetchAllUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func StoreUsers(c echo.Context) error {
	idTemp := c.FormValue(("id"))
	id, err := strconv.Atoi(idTemp)
	nama := c.FormValue(("nama"))
	umurTemp := c.FormValue("umur")
	umur, err := strconv.Atoi(umurTemp)
	address := c.FormValue("alamat")

	result, err := models.StoreUsers(id, nama, umur, address)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)

}

func UpdateUser(c echo.Context) error {
	idTemp := c.FormValue(("id"))
	id, err := strconv.Atoi(idTemp)
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}
	nama := c.FormValue(("nama"))
	umurTemp := c.FormValue("umur")
	umur, err := strconv.Atoi(umurTemp)
	address := c.FormValue("alamat")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.UpdateUser(id, nama, umur, address)
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func DeleteUser(c echo.Context) error {
	idTemp := c.FormValue("id")
	id, err := strconv.Atoi(idTemp)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.DeleteUser(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, result)

}
