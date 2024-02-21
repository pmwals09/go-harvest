package goharvest

import (
	"strings"
	"time"
)

// A wrapper to facilitate marshalling time.Time types into a KitchenTime
// string for API calls, and vice versa. Note that when unmarhsalling
// JSON, we have only the time text string, and not the date, so the date
// portion of the time.Time value will be unreliable.
type KitchenTime struct {
	time.Time
}

func (kitchenTime *KitchenTime) UnmarshalJSON(b []byte) error {
	noQuote := strings.Trim(string(b), `"`)
	if strings.ToLower(noQuote) == "null" {
		kitchenTime = nil
		return nil
	}
	original, err := time.Parse(time.Kitchen, strings.ToUpper(noQuote))
	if err != nil {
		kitchenTime.Time = time.Time{}
		return err
	}
	kitchenTime.Time = original
	return nil
}

func (kitchenTime KitchenTime) MarshalJSON() ([]byte, error) {
	string := kitchenTime.Format(time.Kitchen)
	return []byte(`"` + strings.ToLower(string) + `"`), nil
}
