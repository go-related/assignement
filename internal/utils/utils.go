package utils

import "time"

func ConvertStringToDate(sc string) (time.Time, error) {
	layout2 := time.DateOnly
	t, err := time.Parse(layout2, sc)
	return t.UTC(), err
}

func ConvertDateToString(dt time.Time) string {
	var defaultTime time.Time
	if dt.Equal(defaultTime) {
		return ""
	}
	return dt.UTC().Format(time.DateOnly)
}

func IsDateEmpty(dt *time.Time) bool {
	var defaultTime time.Time
	if dt == nil || (*dt).Equal(defaultTime) {
		return true
	}
	return false
}
