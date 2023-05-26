package numeric

import (
	"math"
	"strconv"
)


func RoundFloat2Point(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

func ConvertStringToFloat(input string) float64 {
	floatRes, _ := strconv.ParseFloat(input, 3)
	return floatRes
}

func ConvertStringToInt(input string) int {
	intRes, _ := strconv.Atoi(input)
	return intRes
}