package uuid

import (
	"fmt"
	"math/rand"
	"time"
)

func GeneratePinCodeUuid() (string, error) {
	rand.Seed(time.Now().UnixNano())
	min := 0
	max := 999999
	pinCode := fmt.Sprintf("%06d", rand.Intn(max-min+1)+min)
	return pinCode, nil
}

func GenerateUuid(prefix string) (string, error) {
	rand.Seed(time.Now().UnixNano())
	min := 0
	max := 9999999999
	pinCode := fmt.Sprintf("%s%06d", prefix, rand.Intn(max-min+1)+min)
	return pinCode, nil
}

func GenerateUuidWithCurrentTimeStamp() string {
	currentTime := time.Now()
	tmFormat := fmt.Sprintf("%02d%02d%02d%02d%02d%02d%02d", currentTime.Year(), currentTime.Month(), currentTime.Day(), currentTime.Hour(), currentTime.Minute(), currentTime.Second(), currentTime.Nanosecond()/1e6)
	return tmFormat
}
