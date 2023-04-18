package resolvers

import "time"

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func parseDateString(dateString string) (time.Time, error) {
	// Layout defines the format of the input string.
	layout := "2006-01-02T15:04:05.999999"
	// Parse the input string into a time.Time value.
	parsedTime, err := time.Parse(layout, dateString)
	if err != nil {
		return time.Time{}, err
	}
	return parsedTime, nil
}
