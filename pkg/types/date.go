package types

import (
	"fmt"
	"time"
)

type DateOnly struct {
	time.Time
}

func (d *DateOnly) ToString() string {
	return d.Format("2006-01-02")
}

func (d *DateOnly) UnmarshalJSON(b []byte) error {
	str := string(b)
	str = str[1 : len(str)-1]
	t, err := time.Parse(time.DateOnly, str)
	if err != nil {
		return err
	}
	d.Time = t
	return nil
}

func (d DateOnly) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", d.ToString())), nil
}
