package service

import (
	"math/rand"
)

const (
	MaxValue = 10000.0
	MinValue = 1.0
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
		return MaxValue
	}

	return MinValue
}
