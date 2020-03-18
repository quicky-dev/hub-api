package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo"
	filestore "github.com/quicky-dev/api/fileStore"
)

/* --------------------------------- GetFile -------------------------------- */

//GetFile takes in uuid and sends user the file to the install via CL
func GetUbuntuFile(c echo.Context) error {
	uuid := c.Param("uuid")
	fmt.Println(uuid)
	// get S3 Handler
	handler, err := filestore.GetHandler()
	if err != nil {
		log.Fatal("There was an error getting S3 Session: ", err)
	}
	script, err := handler.ReadFile("ubuntu", uuid)
	if err != nil {
		fmt.Println("KEY: ", uuid)
		fmt.Println("There was an error getting setup script: ", err)
		return c.JSON(http.StatusBadRequest, err)
	}
	// sends the setup script as string
	return c.String(http.StatusOK, script)
}

/* -------------------------------- GetItems -------------------------------- */

//GetItems sends struct of supported items for download
func GetUbuntuItems(c echo.Context) error {
	// sends all supported mac packages as big JSON obj
	return c.JSON(http.StatusOK, ubuntu.AvailablePackages)
}
