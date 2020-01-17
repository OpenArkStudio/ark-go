package util

import "time"

func GetNowTime() int64 {
	return time.Now().Unix()
}
