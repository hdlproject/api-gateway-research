package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"strconv"
)

type (
	Employee struct {
		ID        int    `json:"id"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Address   string `json:"address"`
		Age       int    `json:"age"`
		CompanyID int    `json:"company_id"`
	}

	httpAPI struct {
		router *gin.Engine
	}
)

func newHTTPAPI() *httpAPI {
	router := gin.Default()
	router.GET("/employee/:id", getEmployee)
	router.GET("/employee", getEmployees)
	router.GET("/company/:id", getCompany)
	router.GET("/company", getCompanies)

	return &httpAPI{
		router: router,
	}
}

func (instance *httpAPI) execute() {
	if err := instance.router.Run(":8081"); err != nil {
		panic(err)
	}
}

func getEmployee(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var employees []Employee
	_ = json.Unmarshal(readEmployee(), &employees)

	c.JSON(200, employees[id-1])
}

func getEmployees(c *gin.Context) {
	var employees []Employee
	_ = json.Unmarshal(readEmployee(), &employees)

	c.JSON(200, employees)
}

func getCompany(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var companies []Company
	_ = json.Unmarshal(readCompany(), &companies)

	c.JSON(200, companies[id-1])
}

func getCompanies(c *gin.Context) {
	var companies []Company
	_ = json.Unmarshal(readCompany(), &companies)

	c.JSON(200, companies)
}
