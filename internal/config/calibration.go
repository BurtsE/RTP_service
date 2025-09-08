package config

import (
	"encoding/json"
	"log"
	"os"
)

type Calibration struct {
	Rtps []float64 `json:"rtps" yaml:"rtps"`
	Ks   []float64 `json:"ks" yaml:"ks"`
}

func NewCalibration(filename string) Calibration {
	c := Calibration{}

	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("could not open calibration file: %v", err)
	}

	json.Unmarshal(data, &c)

	return c
}
