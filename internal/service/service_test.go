package service

import (
	"math"
	"math/rand/v2"
	"testing"
	"time"
)

const (
	Precision         = 0.02    // Точность вычисление rtp для тестов
	MaxValue          = 10000.0 // Предел входных значений
	SequenceMinLength = 10000
)

// Параметры распределения для x. Считаем, что распределение нормальное
var (
	Mean   = (MaxValue + 1) / 2
	StdDev = MaxValue / 6
)

func TestGeneration(t *testing.T) {
	tests := []struct {
		name     string
		sequence []float64
		rtp      float64
	}{
		{
			name:     "test #1",
			sequence: generateNormalDistributedSequence(),
			rtp:      0.1,
		},
		{
			name:     "test #2",
			sequence: generateNormalDistributedSequence(),
			rtp:      0.2,
		},
		{
			name:     "test #3",
			sequence: generateNormalDistributedSequence(),
			rtp:      0.3,
		},
		{
			name:     "test #4",
			sequence: generateNormalDistributedSequence(),
			rtp:      0.4,
		},
		{
			name:     "test #5",
			sequence: generateNormalDistributedSequence(),
			rtp:      0.5,
		},
		{
			name:     "test #6",
			sequence: generateNormalDistributedSequence(),
			rtp:      0.6,
		},
		{
			name:     "test #7",
			sequence: generateNormalDistributedSequence(),
			rtp:      0.7,
		},
		{
			name:     "test #8",
			sequence: generateNormalDistributedSequence(),
			rtp:      0.8,
		},
		{
			name:     "test #9",
			sequence: generateNormalDistributedSequence(),
			rtp:      0.98,
		},
		{
			name:     "test #10",
			sequence: generateNormalDistributedSequence(),
			rtp:      1,
		},
		{
			name:     "test #11",
			sequence: generateConst(5000),
			rtp:      0.1,
		},
		{
			name:     "test #12",
			sequence: generateConst(6000),
			rtp:      0.5,
		},
		{
			name:     "test #13",
			sequence: generateConst(2000),
			rtp:      0.9,
		},
		{
			name:     "test #14",
			sequence: generateConst(1),
			rtp:      1,
		},
		{
			name:     "test #15",
			sequence: generateDefault(),
			rtp:      1,
		},
		{
			name:     "test #16",
			sequence: generateDefault(),
			rtp:      0.5,
		},
		// {
		// 	name:     "test #",
		// 	sequence: generateDefault(10000),
		// 	rtp:      1,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := NewService(tt.rtp)
			transformed := make([]float64, len(tt.sequence))
			sum1 := 0.0
			sum0 := float64(len(tt.sequence))
			totalSum := 0.0

			zeroed := 0
			for i := range tt.sequence {
				totalSum += tt.sequence[i]
				multiplicator := service.GenerateMultiplicator()
				if multiplicator <= tt.sequence[i] {
					transformed[i] = 0
					zeroed++
				} else {
					transformed[i] = tt.sequence[i]
				}
				sum1 += transformed[i]
			}
			t.Log(sum1, sum0, zeroed)

			mean := totalSum / float64(len(tt.sequence))

			estimatedRtp := sum1 / sum0 / mean

			if math.Abs(estimatedRtp-tt.rtp) > Precision {
				t.Logf("sum is not equal to rtp, want: %v, got: %v", tt.rtp, estimatedRtp)
				t.Fail()
			}
		})
	}
}

// generateNormalDistributedSequence генерирует числа
// согласно в соответствии с заданным распределением
// (пакет rand создает равномерное распределение)
func generateNormalDistributedSequence() []float64 {
	v := uint64(time.Now().UnixNano())
	r := rand.New(rand.NewPCG(v, v+3))

	length := SequenceMinLength + r.Int()%3000
	sequence := make([]float64, 0, length)

	for range length {
		sequence = append(sequence, generateNormalFloat64(r, Mean, StdDev))
	}
	return sequence
}

// Преобразование Бокса-Мюллера
func generateNormalFloat64(r *rand.Rand, mean, stdDev float64) float64 {
	u1 := r.Float64()
	u2 := r.Float64()

	z := math.Sqrt(-2*math.Log(u1)) * math.Cos(2*math.Pi*u2)

	return mean + stdDev*z
}

func generateDefault() []float64 {
	v := uint64(time.Now().UnixNano())
	r := rand.New(rand.NewPCG(v, v+3))
	length := SequenceMinLength + r.Int()%3000
	sequence := make([]float64, 0, length)

	for range length {
		val := r.Float64() * MaxValue
		sequence = append(sequence, val)
	}
	return sequence
}

func generateConst(val float64) []float64 {
	v := uint64(time.Now().UnixNano())
	r := rand.New(rand.NewPCG(v, v+3))
	length := SequenceMinLength + r.Int()%3000
	sequence := make([]float64, 0, length)

	for range length {
		sequence = append(sequence, val)
	}
	return sequence
}
