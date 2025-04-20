package enum

import (
	"encoding/json"
	"errors"
)

type Status string

const (
	StatusOpen   Status = "OPEN"
	StatusClosed Status = "CLOSED"
)

func (s Status) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(s))
}

func (s *Status) UnmarshalJSON(b []byte) error {
	var str string
	if err := json.Unmarshal(b, &str); err != nil {
		return err
	}
	switch str {
	case string(StatusOpen):
		*s = StatusOpen
		return nil
	case string(StatusClosed):
		*s = StatusClosed
		return nil
	default:
		return errors.New("invalid")
	}
}
