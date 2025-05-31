package utils

import (
	"strconv"
	"time"
)

func ParseUint(s string) (uint, error) {
	id64, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0, ErrInvalidID
	}
	return uint(id64), nil
}

func ParseInt(s string) (int, error) {
	id64, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, ErrInvalidID
	}
	return int(id64), nil
}

func GetCurrentTime() time.Time {
	return time.Now()
}
