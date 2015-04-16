package cm15

import "time"

// Wrapper around time.Time that adds the ability to unmarshal ruby json date time values
type RubyTime struct {
	time.Time
}

// Implement unmarshaller interface
func (r *RubyTime) Unmarshal(data []byte, v interface{}) error {
	t, err := time.Parse("2006/01/02 15:04:05 -0700", string(data))
	if err != nil {
		return err
	}
	r.Time = t
	return nil
}
