package random

import "math/rand"

type RandomStringService struct {
	randomCharacterPool string
}

func NewService(randomPool string) *RandomStringService {
	return &RandomStringService{randomPool}
}

func (s *RandomStringService) GenerateRandomString(outputLength int) string {
	b := make([]byte, outputLength)
	for i := range b {
		b[i] = s.randomCharacterPool[rand.Intn(len(s.randomCharacterPool))]
	}
	return string(b)
}
