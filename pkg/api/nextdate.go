package api

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// NextDate функция расчитывает следующюю дату в зависимости от правила
func NextDate(now time.Time, dstart string, repeat string) (string, error) {
	parameters := "dywm"
	date, err := time.Parse(FormatDate, dstart)
	if err != nil {
		return "", fmt.Errorf("string Parsing error:%w", err)
	}
	if repeat == "" {
		return "DELETE", nil
	}
	sliceRepeat := strings.Split(repeat, " ")
	if !strings.Contains(parameters, sliceRepeat[0]) {
		return "", errors.New("invalid character")
	}
	if len(sliceRepeat) == 1 && sliceRepeat[0] != "y" {
		return "", errors.New("the interval is not specified")
	}

	switch sliceRepeat[0] {
	case "d":
		interval, err := strconv.Atoi(sliceRepeat[1])
		if err != nil {
			return "", fmt.Errorf("string conversion error:%w", err)
		}
		if interval <= 0 || interval > 400 {
			return "", errors.New("The maximum interval has been exceeded")
		}
		for {
			date = date.AddDate(0, 0, interval)
			if afterNow(date, now) {
				break
			}
		}
		return date.Format(FormatDate), nil
	case "y":
		for {
			date = date.AddDate(1, 0, 0)
			if afterNow(date, now) {
				break
			}
		}
		return date.Format(FormatDate), nil
	case "w":
		return "", errors.New("unsupported format")
	case "m":
		return "", errors.New("unsupported format")
	}
	return "", errors.New("unknown error func NextDate")
}

// AfterNow функция  возвращает true, если первая дата больше второй
func afterNow(date, now time.Time) bool {
	return date.Format(FormatDate) > now.Format(FormatDate)
}
