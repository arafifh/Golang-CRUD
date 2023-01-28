package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Employee struct {
	Id			string
	Name 		string
	Age 		int
	Division 	string
}

var employees = []Employee {
	{Id: "c1", Name: "Rafif", Age: 19, Division: "Backend"},
}

func main() {
	r := gin.Default()
	r.GET("/employees", getEmployees)
	r.GET("/employee/:Id", detailEmployee)
	r.POST("/employee", createEmployee)
	r.PUT("/employee/update/:Id", updateEmployee)
	r.DELETE("/employee/delete/:Id", deleteEmployee)

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
		"message": "404 Not Found",
		})
	})
  	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func getEmployees(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H {
		"status":	"success",
		"data":		employees,
	})
}

func detailEmployee(c *gin.Context) {
	Id := c.Param("Id")
	condition := false
	var detailEmployee Employee

	for i, Employee := range employees {
		if Id == Employee.Id {
			condition = true
			detailEmployee = employees[i]
			break
		}
	}

	if !condition {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H {
			"error":	"Data Not Found",
			"message":	fmt.Sprintf("employee with id %v not found", Id),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H {
		"status":	"success",
		"data":		detailEmployee,
	})
}

func createEmployee(c *gin.Context) {
	var newEmployee Employee
	err := c.ShouldBindJSON(&newEmployee)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	newEmployee.Id = fmt.Sprintf("c%d", len(employees) + 1)
	employees = append(employees, newEmployee)

	c.JSON(http.StatusCreated, gin.H {
		"status":	"success",
		"message":	"new employee created successfully",
		"data":		newEmployee,
	})
}

func updateEmployee(c *gin.Context) {
	Id := c.Param("Id")
	condition := false
	var updateEmployee Employee

	err := c.ShouldBindJSON(&updateEmployee)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H {
			"error":	err.Error(),
		})
		return
	}

	for i, Employee := range employees {
		if Id == Employee.Id {
			condition = true
			employees[i] = updateEmployee
			employees[i].Id = Id
			break
		}
	}

	if !condition {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H {
			"error": "Data Not Found",
			"message": fmt.Sprintf("employee with id %v not found", Id),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H {
		"status":	"success",
		"message":	fmt.Sprintf("employee with id %v updated successfully", Id),
		"data": 	updateEmployee,
	})
}

func deleteEmployee(c *gin.Context) {
	Id := c.Param("Id")
	condition := false
	var employeeIndex int

	for i, Employee := range employees {
		if Id == Employee.Id {
			condition = true
			employeeIndex = i
			break
		}
	}

	if !condition {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H {
			"status":	"Data Not Found",
			"message":	fmt.Sprintf("employee with id %v not found", Id),
		})
		return
	}

	copy(employees[employeeIndex:], employees[employeeIndex+1:])
	employees[len(employees)-1] = Employee{}
	employees = employees[:len(employees)-1]

	c.JSON(http.StatusOK, gin.H {
		"status":	"success",
		"message":	fmt.Sprintf("employee with id %v deleted successfully", Id),
	})
}