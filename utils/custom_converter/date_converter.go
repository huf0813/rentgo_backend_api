package custom_converter

import (
	"time"
)

func NewDateConverter(val string) (time.Time, error) {
	layout := "2006-01-02T15:04:05.000Z"
	t, err := time.Parse(layout, val)

	if err != nil {
		return t, err
	}
	return t, nil
}
