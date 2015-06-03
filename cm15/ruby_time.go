package cm15 // import "gopkg.in/rightscale/rsc.v2/cm15"

import "time"

// Wrapper around time.Time that adds the ability to unmarshal ruby json date time values
type RubyTime struct {
	time.Time
}

// Implement unmarshaller interface
func (r *RubyTime) UnmarshalJSON(b []byte) (err error) {
	s := string(b)
	t, err := time.Parse("2006/01/02 15:04:05 -0700", s[1:len(s)-1])
	if err != nil {
		return err
	}
	r.Time = t
	return nil
}
