package cm15

import (
	"encoding/json"
	"time"
)

// RubyTime is a wrapper around time.Time that adds the ability to unmarshal ruby JSON date time
// values.
type RubyTime struct {
	time.Time
}

// MarshalJSON implements the marshaller interface.
func (r *RubyTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.Format("2006/01/02 15:04:05 -0700"))
}

// UnmarshalJSON implements the unmarshaller interface.
func (r *RubyTime) UnmarshalJSON(b []byte) (err error) {
	s := string(b)
	t, err := time.Parse("2006/01/02 15:04:05 -0700", s[1:len(s)-1])
	if err != nil {
		return err
	}
	r.Time = t
	return nil
}
