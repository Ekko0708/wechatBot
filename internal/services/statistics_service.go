package services

import (
	"sync"
)

type StatisticsService struct {
	mu    sync.Mutex
	count int
}

func (s *StatisticsService) Increment() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.count++
}

func (s *StatisticsService) GetCount() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.count
}
