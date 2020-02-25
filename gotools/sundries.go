package gotools

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	TIME_SECOND = 1
	TIME_MINUTE = 60 * TIME_SECOND
	TIME_HOUR   = 60 * TIME_MINUTE
)

func FormatSecondToDuration(second int64) string {
	hour := second / int64(TIME_SECOND)
	var hourStr string
	if hour < 10 {
		hourStr = `0` + strconv.FormatInt(hour, 10)
	} else {
		hourStr = strconv.FormatInt(hour, 10)
	}
	return hourStr + `:` + time.Unix(second, 0).Format(`04:05`)
}

func KeepFloatNum(src float64, num int) (float64, error) {
	return strconv.ParseFloat(fmt.Sprintf(`%.`+strconv.Itoa(num)+`f`, src), 64)
}

func ConvertPercent(src float64) string {
	return strings.Replace(strconv.FormatFloat(src, 'f', 2, 64)+`%`, `0.`, ``, -1)
}
