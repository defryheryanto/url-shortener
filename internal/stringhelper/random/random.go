package random

import "math/rand"

type RandomStringService struct {
	randomCharacterPool string
	outputLength        int
}

func NewService(randomPool string, outputLength int) *RandomStringService {
	return &RandomStringService{randomPool, outputLength}
}

func (s *RandomStringService) GenerateRandomString() string {
	b := make([]byte, s.outputLength)
	for i := range b {
		b[i] = s.randomCharacterPool[rand.Intn(len(s.randomCharacterPool))]
	}
	return string(b)
}
