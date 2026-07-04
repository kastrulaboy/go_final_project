package api

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func afterNow(date, now time.Time) bool {
	return date.After(now)
}

func parseDate(date string) (time.Time, error) {
	return time.Parse("20060102", date)
}

func NextDate(now time.Time, dstart string, repeat string) (string, error) {
	if repeat == "" {
		return "", fmt.Errorf("не указано правило повторения")
	}

	date, err := parseDate(dstart)
	if err != nil {
		return "", fmt.Errorf("неверный формат даты: %w", err)
	}

	parse := strings.Fields(repeat)

	if len(parse) == 0 {
		return "", fmt.Errorf("неверный формат правила повторения")
	}

	switch parse[0] {
	case "d":
		if len(parse) != 2 {
			return "", fmt.Errorf("правило 'd' должно иметь вид: d <число>")
		}

		interval, err := strconv.Atoi(parse[1])
		if err != nil {
			return "", fmt.Errorf("интервал должен быть целым числом")
		}

		if interval < 1 || interval > 400 {
			return "", fmt.Errorf("интервал должен быть в диапазоне от 1 до 400")
		}

		for {
			date = date.AddDate(0, 0, interval)

			if afterNow(date, now) {
				break
			}
		}

	case "y":
		if len(parse) != 1 {
			return "", fmt.Errorf("правило 'y' не принимает параметров")
		}

		for {
			date = date.AddDate(1, 0, 0)

			if afterNow(date, now) {
				break
			}
		}

	default:
		return "", fmt.Errorf("неизвестное правило повторения")
	}

	return date.Format("20060102"), nil
}

func nextDayHandler(res http.ResponseWriter, req *http.Request) {
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
			http.Error(res, "неверный формат даты now, ожидается YYYYMMDD", http.StatusBadRequest)
			return
		}
	}

	next, err := NextDate(now, date, repeat)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	res.Write([]byte(next))
}