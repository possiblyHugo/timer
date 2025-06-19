package main

import (
	"fmt"
	"time"
)

type SecondsTimer struct {
	timer     *time.Timer
	end       time.Time
	remaining time.Duration
}

func NewSecondsTimer(t time.Duration) (*SecondsTimer, error) {
	if t <= 0 {
		return nil, fmt.Errorf("Invalid Duration: Must be greater than zero")
	}
	return &SecondsTimer{time.NewTimer(t), time.Now().Add(t), time.Second}, nil
}

func (s *SecondsTimer) Reset(t time.Duration) {
	s.timer.Reset(t)
	s.end = time.Now().Add(t)
}

func (s *SecondsTimer) Stop() {
	s.timer.Stop()
	s.remaining = s.TimeRemaining() // Save time for later use in the Resume Method
}

func (s *SecondsTimer) Resume() {
	s.timer = time.NewTimer(s.remaining)
}

func (s *SecondsTimer) TimeRemaining() time.Duration {
	return s.end.Sub(time.Now())
}
