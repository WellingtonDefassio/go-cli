package utils

import (
	"strconv"
	"time"
)

func FormatEpoc(epoc string) string {
	epoch, err := strconv.ParseInt(epoc, 10, 64)
	if err != nil {
		return "error to convert epoc to int"
	}
	unixTime := time.Unix(epoch, 0)
	return unixTime.Format("15:04:05 02/01/2006")
}
