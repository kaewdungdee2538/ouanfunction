package formatString

import (
	"fmt"
	"time"
	"strings"
)

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

func StringToCamelCase(inputString string) string {
	words := strings.Fields(inputString)
	if len(words) == 0 {
		return ""
	}

	camelCase := strings.ToLower(words[0])
	for _, word := range words[1:] {
		if len(word) > 0 {
			wordLower := strings.ToLower(word)
			camelCase += strings.Title(wordLower)
		}
	}

	return camelCase
}

func RemoveWhiteSpace(input string) string {
	cleanedString := strings.Replace(input, "\u200B", "", -1)
	return cleanedString
}


func SetValueForWhereLikeData(input string) string {
	likeChar := "%"
	result := fmt.Sprintf("%s%s%s", likeChar, input, likeChar)
	return result
}


func ConvertArrayToString(a []int, delim string) string {
    return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
    //return strings.Trim(strings.Join(strings.Split(fmt.Sprint(a), " "), delim), "[]")
    //return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(a)), delim), "[]")
}
