package utils

import (
	"strings"
	"time"
)

type JsonDate time.Time
type JsonTime time.Time

func unmarshal(b []byte, layout *string) (*time.Time, error) {

	value := strings.Trim(string(b), `"`) //get rid of "
	if value == "" || value == "null" {
		return nil, nil
	}

	t, err := time.Parse(*layout, value) //parse time
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func (c *JsonTime) UnmarshalJSON(b []byte) error {
	layout := "2006-01-02 15:04:05"

	if t, err := unmarshal(b, &layout); err != nil {
		return err
	} else {
		*c = JsonTime(*t) //set result using the pointer
	}
	return nil
}

func (c *JsonTime) MarshalJSON() ([]byte, error) {
	return []byte(`"` + time.Time(*c).Format("2006-01-02 15:04:05") + `"`), nil
}

func (c *JsonDate) UnmarshalJSON(b []byte) error {
	layout := "2006-01-02"

	if t, err := unmarshal(b, &layout); err != nil {
		return err
	} else {
		*c = JsonDate(*t) //set result using the pointer
	}
	return nil
}

func (c *JsonDate) MarshalJSON() ([]byte, error) {
	return []byte(`"` + time.Time(*c).Format("2006-01-02") + `"`), nil
}
