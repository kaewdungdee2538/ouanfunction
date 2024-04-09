package time

import (
	"fmt"
	"math"
	"strings"
	"time"
)

func ConvertDatetimeToString(timeInput time.Time) string {
	return timeInput.Format("2006-01-02 15:04:05")
}

func ConvertDatetimeToDateString(timeInput time.Time) string {
	return timeInput.Format("2006-01-02")
}

func ConvertStringToNewDateTimeFormat(str string) string {

	if newTime, err := ConvertStringToDateTimeAndError(str); err == nil {
		return newTime.Format("02/01/2006 15:04:05")
	} else {
		return str
	}

}

func ConvertStringToDateTimeAndError(str string) (time.Time, error) {
	layout := "2006-01-02 15:04:05"
	t, err := time.Parse(layout, str)
	if err != nil {
		layout2 := "2006-01-02 15:04:05"
		t2, err := time.Parse(layout2, str)
		if err != nil {
			fmt.Println(err)
			return time.Now(), err
		}
		return t2, nil
	}
	return t, nil
}

func ConvertStringToDateTime(str string) time.Time {
	layout := "2006-01-02 15:04:05"
	t, err := time.Parse(layout, str)
	if err != nil {
		layout2 := "2006-01-02 15:04:05"
		t2, err := time.Parse(layout2, str)
		if err != nil {
			fmt.Println(err)
			return time.Now()
		}
		return t2
	}
	return t
}

func ConvertStringToTime(str string) time.Time {
	layout := "15:04:05"
	t, err := time.Parse(layout, str)
	if err != nil {
		layout2 := "15:04:05"
		t2, err := time.Parse(layout2, str)
		if err != nil {
			fmt.Println(err)
			return time.Now()
		}
		return t2
	}
	return t
}

func CalculateTimeToMinutes(timeStartInput string, timeStopInput string) float64 {
	timeStart := ConvertStringToTime(timeStartInput)
	timeStop := ConvertStringToTime(timeStopInput)
	timeDiff := timeStop.Sub(timeStart)

	return timeDiff.Minutes()
}

func ConvertDatetimeToDate(timeInput time.Time) time.Time {
	date := fmt.Sprintf("%v-%02d-%02d 00:00:00", timeInput.Year(), timeInput.Month(), timeInput.Day())
	return ConvertStringToDateTime(date)
}

func ConvertDatetimeToTime(timeInput time.Time) time.Time {
	tm := fmt.Sprintf("%02d:%02d:%02d", timeInput.Hour(), timeInput.Minute(), timeInput.Second())
	return ConvertStringToTime(tm)
}

func CalculateBetweenTimeInAndTimeOutToDuration(timeIn time.Time, timeOut time.Time) time.Duration {
	timeDiff := timeOut.Sub(timeIn)
	return timeDiff
}

func ConvertIntToDuration(durationInt int) time.Duration {
	durationOutput := time.Duration(durationInt * int(time.Second))
	return durationOutput
}

func ConverDurationToText(durationInput time.Duration) string {
	minutes := durationInput.Minutes()
	hours := minutes / 60
	days := hours / 24

	dayOut := math.Floor(days)
	hourOut := math.Floor(math.Mod(hours, 24))
	minuteOut := math.Ceil(math.Mod(durationInput.Minutes(), 60))

	if minuteOut >= 60 {
		minuteOut = 0
		hourOut++
	}
	if hourOut >= 24 {
		hourOut = 0
		dayOut++
	}

	if dayOut > 0 {
		return fmt.Sprintf(`%d วัน %d ชั่วโมง %d นาที`, int(dayOut), int(hourOut), int(minuteOut))
	} else if hourOut > 0 {
		return fmt.Sprintf(`%d ชั่วโมง %d นาที`, int(hourOut), int(minuteOut))
	} else {
		return fmt.Sprintf(`%d นาที`, int(minuteOut))
	}
}

func ConverStringToDurationText(durationText string) string {
	if len(durationText) == 0 {
		return ""
	}
	_time := ConvertStringToTime(durationText)

	hourOut := _time.Hour()
	minuteOut := _time.Minute()

	if minuteOut >= 60 {
		minuteOut = 0
		hourOut++
	}
	if hourOut >= 24 {
		hourOut = 0
	}

	if hourOut > 0 {
		return fmt.Sprintf(`%d ชั่วโมง %d นาที`, int(hourOut), int(minuteOut))
	} else {
		return fmt.Sprintf(`%d นาที`, int(minuteOut))
	}
}

func GenerateTimeUnitText(timeMinutes int) interface{} {
	if timeMinutes <= 0 {
		return nil
	} else {
		duration := time.Duration(timeMinutes) * time.Minute

		return fmt.Sprintf(`จอดฟรี %s`, ConverDurationToText(duration))
	}
}

func GenerateCurrencyUnitText(currencyUnit int) interface{} {
	if currencyUnit <= 0 {
		return nil
	} else {

		return fmt.Sprintf(`ส่วนลด %d บาท`, currencyUnit)
	}
}

func CheckDateIsWeekend(dateInput time.Time) bool {
	weekDayNumber := dateInput.Weekday()
	if weekDayNumber == 0 || weekDayNumber == 6 {
		return true
	}
	return false
}

func CalParkingInterval(parkingInTime string, parkingOutTime string) string {
	parkingTimeIn := ConvertStringToDateTime(parkingInTime)
	parkingTimeOut := ConvertStringToDateTime(parkingOutTime)
	minutesTimeDiff := CalculateBetweenTimeInAndTimeOutToDuration(parkingTimeIn, parkingTimeOut)
	return ConverDurationToText(minutesTimeDiff)
}

func ConvertStringDatetime(str string) string {

	layout := "2006-01-02T15:04:05.000000"
	t, err := time.Parse(layout, str)
	if err != nil {
		layout2 := "2006-01-02T15:04:05.000Z"
		t2, err := time.Parse(layout2, str)
		if err != nil {
			fmt.Println(err)
			return str
		}
		return t2.Format("2006-01-02 15:04:05")
	}
	timeText := t.Format("2006-01-02 15:04:05")
	return timeText
}

func SplitTime(str string) (string, string, string) {
	s := strings.Split(str, ":")
	fmt.Printf("Split Time => %s\n", s)
	if len(s) == 3 {
		return s[0], s[1], s[2]
	} else {
		return "00", "00", "00"
	}
}

func ConvertSecondsToMinutes(sec int) int {
	return int(math.Ceil(float64(sec) / 60))
}

func ConvertMinutesToHours(sec int) int {
	return int(math.Ceil(float64(sec) / 60))
}

func GetUnixTimestampText() string {
	unixText := fmt.Sprint(time.Now().UnixMilli())
	return unixText
}


func ConvertUnixTimestampToDateTime(unixTimestamp int64) string {
	timestamp := time.Unix(unixTimestamp, 0)
	return timestamp.Format("2006-01-02 15:04:05")
}
