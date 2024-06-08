package library

const secondInMinute int = 60
const minuteInHour int = 60
const DayInHour int = 24

func secondToMinute(sec int) int {
	return sec / secondInMinute
}

func MinuteToHour(min int) int{
	return min / minuteInHour
}