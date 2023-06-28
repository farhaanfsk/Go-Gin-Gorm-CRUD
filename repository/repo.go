package repository

import (
	M "Go-Gin-Gorm-CRUD/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func (r *Repository) GetEmployees() []M.Employee {
	var employees []M.Employee
	r.DB.Find(&employees)
	return employees
}

func (r *Repository) GetEmployeeById(id uuid.UUID) M.Employee {
	var employee M.Employee
	r.DB.Find(&employee)
	return employee
}
