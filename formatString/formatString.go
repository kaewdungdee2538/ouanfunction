package formatString

import (
	"fmt"
	"time"
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