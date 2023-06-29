package main

import (
	"Go-Gin-Gorm-CRUD/models"
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func main() {
	r := gin.Default()
	r.GET("/employees", getEmployees)
	r.GET("/employees/:id", getEmployeesById)
	r.POST("/employees", CreateEmployee)
	r.PUT("/employees", UpdateEmployee)
	r.DELETE("/employees/:id", deleteEmployee)
	r.Run()
}

func UpdateEmployee(c *gin.Context) {
	var emp models.Employee
	err := c.Bind(&emp)
	if err != nil {
		fmt.Println("unable to bind employee", err)
	}
	var existingEmp models.Employee
	result := DB.Preload("Address").First(&existingEmp, "id = ?", emp.Id)
	if result.Error != nil {
		fmt.Println(result.Error)
		c.JSON(http.StatusNotFound, gin.H{"data": "employee not found"})
		return
	}
	emp.Id = existingEmp.Id
	result = DB.Updates(&emp)
	if result.Error != nil {
		fmt.Println(result.Error)
		c.JSON(http.StatusNotFound, gin.H{"data": "employee not found"})
		return
	}
	c.JSON(http.StatusOK, emp)
}

func deleteEmployee(c *gin.Context) {
	id := c.Param("id")
	result := DB.Delete(&models.Employee{}, "id = ?", id).Delete(&models.Address{})
	if result.Error != nil {
		fmt.Println(result.Error)
		c.JSON(http.StatusBadRequest, gin.H{"data": "Unable to delete employee"})
		return
	}
	c.JSON(http.StatusOK, gin.H{})

}

func getEmployees(c *gin.Context) {
	var employees []models.Employee
	result := DB.Preload("Address").Find(&employees)
	fmt.Println(employees)
	if result.Error != nil {
		fmt.Println(result.Error)
		c.JSON(http.StatusNotFound, gin.H{"data": "employee not found"})
		return
	}
	c.JSON(http.StatusOK, employees)
}

var DB *gorm.DB

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("unable to load the env file", err)
	}
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"), os.Getenv("DB_SSLMODE"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("unable to connect to db", err)
	}
	db.AutoMigrate(models.Employee{}, models.Address{})
	DB = db
}

func CreateEmployee(c *gin.Context) {
	var emp models.Employee
	err := c.Bind(&emp)
	if err != nil {
		fmt.Println("unable to bind employee", err)
	}
	emp.Id = uuid.New()
	emp.Address.Id = uuid.New()
	DB.Create(&emp)
	c.JSON(http.StatusOK, emp)
}

func getEmployeesById(c *gin.Context) {
	id := c.Param("id")
	var emp models.Employee
	result := DB.Preload("Address").First(&emp, "id = ?", id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"data": "employee not found"})
		return
	}
	c.JSON(http.StatusOK, emp)
}
