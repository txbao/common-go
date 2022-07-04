package holidays

import "time"

func getDate(value string) (time.Time, error) {
	valueTime, err := time.Parse("2006/01/02", value)
	if err != nil {
		return time.Time{}, err
	}
	return valueTime, nil
}
