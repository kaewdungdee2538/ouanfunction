package validation

import (
	"fmt"
	"net"
	"net/mail"
	"regexp"
	"strings"
	"time"

	_time "github.com/kaewdungdee2538/ouanfunction/v2/time"
)

func IsNotStringAlphabet(str string) bool {
	const alpha = `/!@#$%^&*()_+\-=\[\]{};':"|,.<>\/?~`
	for _, char := range str {
		if strings.Contains(alpha, strings.ToLower(string(char))) {
			return true
		}
	}
	return false
}

func IsNotStringAlphabetRemark(str string) bool {
	const alpha = `!@#$%^&*\[\]{};':",<>?~|`
	for _, char := range str {
		if strings.Contains(alpha, strings.ToLower(string(char))) {
			return true
		}
	}
	return false
}

func IsNotStringAlphabetForJSONString(str string) bool {
	const alpha = `!@#$%^*;<>?~|`
	for _, char := range str {
		if strings.Contains(alpha, strings.ToLower(string(char))) {
			return true
		}
	}
	return false
}

func IsNotStringAlphabetForLine(str string) bool {
	const alpha = `-!#$%^&*\[\]{};':",<>?~|`
	for _, char := range str {
		if strings.Contains(alpha, strings.ToLower(string(char))) {
			return true
		}
	}
	return false
}

func IsNotStringAlphabetForEmail(str string) bool {
	const alpha = `!#$%^&*\[\]{};':",<>?~|`
	for _, char := range str {
		if strings.Contains(alpha, strings.ToLower(string(char))) {
			return true
		}
	}
	return false
}

func IsNotStringAlphabetForPassword(str string) bool {
	const alpha = `-#$%^&*\[\]{};':",<>?~`
	for _, char := range str {
		if strings.Contains(alpha, strings.ToLower(string(char))) {
			return true
		}
	}
	return false
}

func IsNotStringAlphabetForMessageString(str string) bool {
	const alpha = `;`
	for _, char := range str {
		if strings.Contains(alpha, strings.ToLower(string(char))) {
			return true
		}
	}
	return false
}

func IsNotStringEngOrNumber(str string) bool {
	for _, charVariable := range str {
		if (charVariable < 'a' || charVariable > 'z') && (charVariable < 'A' || charVariable > 'Z') && (charVariable < '0' || charVariable > '9') {
			return true
		}
	}
	return false
}

func IsNotStringNumber(str string) bool {
	for _, charVariable := range str {
		if charVariable < '0' || charVariable > '9' {
			return true
		}
	}
	return false
}

func IsNotStringFloat(str string) bool {
	for _, charVariable := range str {
		if (charVariable < '0' || charVariable > '9') && charVariable != '.' {
			return true
		}
	}
	return false

}

func IsNotEmailFormat(str string) bool {
	_, err := mail.ParseAddress(str)
	return err != nil
}

func IsNotFormatYear(str string) bool {
	if _, err := time.Parse("2006", str); err != nil {
		if _, err := time.Parse("2006", str); err != nil {
			return true
		}
		return false
	}
	return false
}

func IsNotFormatMonth(str string) bool {
	if _, err := time.Parse("2006-01", str); err != nil {
		if _, err := time.Parse("2006-01", str); err != nil {
			return true
		}
		return false
	}
	return false
}

func IsNotFormatDate(str string) bool {
	if _, err := time.Parse("2006-01-02", str); err != nil {
		if _, err := time.Parse("2006-01-02", str); err != nil {
			return true
		}
		return false
	}
	return false
}

func IsNotFormatDateTime(str string) bool {
	if _, err := time.Parse("2006-01-02 15:04:00", str); err != nil {
		if _, err := time.Parse("2006-01-02 15:04:05", str); err != nil {
			return true
		}
		return false
	}
	return false
}

func IsNotFormatTime(str string) bool {
	if _, err := time.Parse("15:04:00", str); err != nil {
		if _, err := time.Parse("15:04:05", str); err != nil {
			return true
		}
		return false
	}
	return false
}

func IsDatetartAfterTimeEnd(str_start string, str_end string) bool {
	time_start, _ := time.Parse("2006-01-02", str_start)
	time_end, _ := time.Parse("2006-01-02", str_end)
	return time_start.After(time_end)
}

func IsTimeStartAfterTimeEnd(str_start string, str_end string) bool {
	time_start, _ := time.Parse("15:04:05", str_start)
	time_end, _ := time.Parse("15:04:05", str_end)
	return time_start.After(time_end)
}

func IsEmailValid(email string) bool {
	fmt.Print("email format validation : ", email)
	_, err := mail.ParseAddress(email)
	return err == nil
}

func IsDateTimeStartAfterDateTimeEnd(str_start string, str_end string) bool {
	datetime_start, _ := time.Parse("2006-01-02 15:04:05", str_start)
	datetime_end, _ := time.Parse("2006-01-02 15:04:05", str_end)
	return datetime_start.After(datetime_end)
}

func IsDateTimeStartAndDateTimeEndDiffOver31Days(str_start string, str_end string) bool {
	dateStart := _time.ConvertStringToDateTime(str_start)
	dateEnd := _time.ConvertStringToDateTime(str_end)
	subtract := dateEnd.Sub(dateStart)
	hours := subtract.Hours()
	days := hours / 24
	if days > 32 {
		return true
	}
	return false
}

func IsDateTimeStartAndDateTimeEndDiffOver3Months(str_start string, str_end string) bool {
	dateStart := _time.ConvertStringToDateTime(str_start)
	dateEnd := _time.ConvertStringToDateTime(str_end)
	subtract := dateEnd.Sub(dateStart)
	hours := subtract.Hours()
	days := hours / 24
	if days > 93 {
		return true
	}
	return false
}

func IsDateTimeStartAndDateTimeEndDiffOver6Months(str_start string, str_end string) bool {
	dateStart := _time.ConvertStringToDateTime(str_start)
	dateEnd := _time.ConvertStringToDateTime(str_end)
	subtract := dateEnd.Sub(dateStart)
	hours := subtract.Hours()
	days := hours / 24
	if days > 186 {
		return true
	}
	return false
}

func IsDateTimeStartAndDateTimeEndDiffOver1Year(str_start string, str_end string) bool {
	dateStart := _time.ConvertStringToDateTime(str_start)
	dateEnd := _time.ConvertStringToDateTime(str_end)
	subtract := dateEnd.Sub(dateStart)
	hours := subtract.Hours()
	days := hours / 24
	if days > 366 {
		return true
	}
	return false
}

func IsValidIPAddress(ipAddress string) bool {
	// ParseIP returns the IP address (IPv4 or IPv6) represented by the string ipAddress
	// If ipAddress is not a valid IP address, ParseIP returns nil
	return net.ParseIP(ipAddress) != nil
}

func ValidateIntervalTimeFormat(text string) bool {
	regex := regexp.MustCompile("^[0-9]{2,9}:[0-5]{1}[0-9]{1}:[0-5]{1}[0-9]{1}$")
	return regex.MatchString(text)
}
