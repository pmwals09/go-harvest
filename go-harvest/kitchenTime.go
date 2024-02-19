package goharvest

import (
	"strings"
	"time"
)

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
