package service

import (
	"log"
	"math"
	"math/rand"
	"multiplicator/internal/config"
)

type Service struct {
	rtp float64
	k   float64
}

func NewService(rtp float64, calibration config.Calibration) *Service {
	k := findNearestK(calibration, rtp)
	log.Printf("k value set to: %v", k)
	return &Service{
		rtp: rtp,
		k:   k,
	}
}
// GenerateMultiplicator генерирует мультипликаторы, смещая распределение
func (s *Service) GenerateMultiplicator() float64 {
	return 1 + 9999*(math.Pow(rand.Float64(), s.k))
}

func findNearestK(calibration config.Calibration, rtp float64) float64 {
	minDiff := math.MaxFloat64
	nearestK := 0.0
	for i, val := range calibration.Rtps {
		diff := math.Abs(rtp - val)
		if diff < minDiff {
			minDiff = diff
			nearestK = calibration.Ks[i]
		}
	}
	return nearestK
}
