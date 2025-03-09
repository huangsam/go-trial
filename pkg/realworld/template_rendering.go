package realworld

import (
	"embed"
	"text/template"

	"github.com/huangsam/go-trial/internal/util"
)

//go:embed templates/*
var templatesFS embed.FS

// Car is a template-based model for cars.
type Car struct {
	Make       string // Car make
	Model      string // Car model
	ModelYear  int    // Year that the car model was made
	WheelCount int    // Number of wheels
	MileCount  int    // Number of miles driven
}

// Employee is a template-based model for employees.
type Employee struct {
	FirstName string
	LastName  string
	Age       int
	IsManager bool
	Skills    []string
	Salary    int
}

// RenderCarInfo returns info for a car.
//
// The embedded template showcases simple template syntax.
func RenderCarInfo(car *Car) (string, error) {
	tmpl, err := template.ParseFS(templatesFS, "templates/car.template")
	if err != nil {
		return "", err
	}
	return util.RenderToString(tmpl, car)
}

// RenderEmployeeInfo returns info for an employee.
//
// The embedded template showcases complex template syntax.
func RenderEmployeeInfo(emp *Employee) (string, error) {
	tmpl, err := template.ParseFS(templatesFS, "templates/employee.template")
	if err != nil {
		return "", err
	}
	return util.RenderToString(tmpl, emp)
}
