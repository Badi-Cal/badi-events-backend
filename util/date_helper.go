package util

type FirstDayOfWeek int

const (
	Sunday FirstDayOfWeek = 0
	Monday                = 1
)

var firstDayOfWeekMap = map[FirstDayOfWeek]string{
	0: "Sunday",
	1: "Monday",
}

func FirstDayOfWeekString(dayOfWeek FirstDayOfWeek) string {
	return firstDayOfWeekMap[dayOfWeek]
}

func FirstDayOfWeekFromString(str string) FirstDayOfWeek {
	for key, val := range firstDayOfWeekMap {
		if val == str {
			return key
		}
	}

	return Sunday
}
