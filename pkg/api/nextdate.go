package api

import (
	"fmt"
	"strings"
	"strconv"
	"time"
	"net/http"
)

func afterNow(date, now time.Time) bool {
  return date.After(now)
}

func parseDate(date string) (time.Time, error) {
	return time.Parse("20060102", date)
}

func NextDate(now time.Time, dstart string, repeat string) (string, error) {

	if repeat == "" {
		return "", fmt.Errorf("пусто")
	}

	date, err := parseDate(dstart)
	  if err != nil {
		return "", err
	  }

	parse := strings.Fields(repeat)

	if len(parse) == 0 {
		return "", fmt.Errorf("ошибка")
	}

	switch parse[0] {
	case "d":
		if len(parse) != 2 {
			return "", fmt.Errorf("ошибка")
		}

		interval, err := strconv.Atoi(parse[1])
		if err != nil {
			return "", err
		}

		if interval < 1 || interval > 400 {
			return "", fmt.Errorf("ошибка")
		}

		for {
			date = date.AddDate(0, 0, interval)

			if afterNow(date, now) {
				break
			}
		}

	case "y":
		if len(parse) != 1{
			return "", fmt.Errorf("ошибка")
		}

		for {
			date = date.AddDate(1, 0, 0)

			if afterNow(date, now) {
				break
			}
		}
	default:
		return "", fmt.Errorf("ошибка")
		
	}

		return date.Format("20060102"), nil

}

func nextDayHandler( res http.ResponseWriter, req *http.Request) {
	nowStr := req.FormValue("now")
	date := req.FormValue("date")
	repeat := req.FormValue("repeat")

	var now time.Time

	if nowStr == "" {
		now = time.Now()
	} else {

		var err error

		now, err = time.Parse("20060102", nowStr)
		if err != nil {
			http.Error(res, err.Error(), http.StatusBadRequest)
			return
		}
	}

	next, err := NextDate( now, date, repeat)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	res.Write([]byte(next))

}