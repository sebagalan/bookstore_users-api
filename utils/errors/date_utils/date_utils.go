package date_utils

import (
	"time"
)

const (
	apiDateFormat = "2006-01-02T15:04:05Z"
)

//GetNow ...
func GetNow() time.Time {
	return time.Now().UTC()
}

//GetNowSrting ..
func GetNowSrting() string {
	return GetNow().Format(apiDateFormat)
}
