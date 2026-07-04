package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"
	"log"

	"main/pkg/db"
)

func checkDate(task *db.Task) error {
	now := time.Now()

	if task.Date == "" {
		task.Date = now.Format("20060102")
	}

	if _, err := time.Parse("20060102", task.Date); err != nil {
		return err
	}

	var next string

	if task.Repeat != "" {
		var err error
		next, err = NextDate(now, task.Date, task.Repeat)
		if err != nil {
			return err
		}
	}

	today := now.Format("20060102")

	if task.Date < today {
		if task.Repeat == "" {
			task.Date = today
		} else {
			task.Date = next
		}
	}

	return nil
}

func addTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task db.Task

	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		log.Println(err)
		writeJSON(w, map[string]string{
			"error": err.Error(),
		}, http.StatusBadRequest)
		return
	}

	if task.Title == "" {
		writeJSON(w, map[string]string{
			"error": "не указан заголовок",
		}, http.StatusBadRequest)
		return
	}

	err = checkDate(&task)
	if err != nil {
		log.Println(err)
		writeJSON(w, map[string]string{
			"error": err.Error(),
		}, http.StatusBadRequest)
		return
	}

	id, err := db.AddTask(&task)
	if err != nil {
		log.Println(err)
		writeJSON(w, map[string]string{
			"error": err.Error(),
		}, http.StatusInternalServerError)
		return
	}

	writeJSON(w, map[string]string{
		"id": strconv.FormatInt(id, 10),
	}, http.StatusOK)
}