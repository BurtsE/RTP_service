package service

import (
	"math/rand"
)

type Service struct {
	rtp float64
	k   float64
}

func NewService(rtp float64) *Service {
	return &Service{
		rtp: rtp,
	}
}

// GenerateMultiplicator генерирует мультипликаторы

func (s *Service) GenerateMultiplicator() float64 {
	u := rand.Float64()
	if u < s.rtp {
		return 10000
	}

	return 1
}
