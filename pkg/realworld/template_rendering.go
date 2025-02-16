package realworld

import (
	"text/template"

	"github.com/huangsam/go-trial/internal/util"
)

// Car model
type Car struct {
	Make       string // Car make
	Model      string // Car model
	ModelYear  int    // Year that the car model was made
	WheelCount int    // Number of wheels
	MileCount  int    // Number of miles driven
}

// GetCarSimpleInfo returns simple info for a car.
func GetCarSimpleInfo(car *Car) (string, error) {
	tmpl, err := template.ParseFiles(TemplatePath + "/car-simple.template")
	if err != nil {
		return "", err
	}
	return util.RenderToString(tmpl, car)
}
