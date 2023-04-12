package tools

import (
	"strconv"
	"time"
)

func GetTimestamp() int64 {
	return time.Now().UnixNano()
}

func GetStringTimestamp() string {
	return strconv.FormatInt(GetTimestamp(), 10)
}
