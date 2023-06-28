package main

import (
	"Go-Gin-Gorm-CRUD/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func main() {
	r := gin.Default()
	employees = append(employees, models.Employee{
		Id:   uuid.New(),
		Name: "Test",
		Address: models.Address{
			Id:      uuid.New(),
			City:    "hyd",
			State:   "ap",
			Country: "Ind",
		},
	})
	r.GET("/employees", getEmployees)
	r.Run()
}

func getEmployees(c *gin.Context) {
	c.JSON(http.StatusOK, employees)
}

var employees []models.Employee
