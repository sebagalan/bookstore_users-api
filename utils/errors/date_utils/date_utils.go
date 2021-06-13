package date_utils

import (
	"time"
)

//APIDateFormat ...
const (
	APIDateFormat = "2006-01-02T15:04:05Z"
	APIDBLayout   = "2006-01-02 15:04:05"
)

//GetNow ...
func GetNow() time.Time {
	return time.Now().UTC()
}

//GetNowSrting ..
func GetNowSrting() string {
	return GetNow().Format(APIDateFormat)
}

//GetNowSrting ..
func GetNowDBSrting() string {
	return GetNow().Format(APIDBLayout)
}
