package goharvest

import (
	"time"
)

type Date struct {
	time.Time
}

func (s *Date) UnmarshalJSON(input []byte) error {
	newTime, err := time.Parse(time.DateOnly, string(input[1:len(input)-1]))
	if err != nil {
		s.Time = time.Time{}
		return err
	}

	s.Time = newTime
	return nil
}

func (s Date) MarshalJSON() ([]byte, error) {
	str := s.Format(time.DateOnly)
	return []byte(`"` + str + `"`), nil
}
