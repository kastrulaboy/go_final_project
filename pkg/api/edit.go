package api

import (
	"encoding/json"
	"net/http"

	"main/pkg/db"
)

func editTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task db.Task

	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		writeJSON(w, map[string]string{
			"error": err.Error(),
		})
		return
	}

	if task.ID == "" {
		writeJSON(w, map[string]string{
			"error": "ошибка",
		})
		return
	}

	if task.Title == "" {
		writeJSON(w, map[string]string{
			"error": "ошибка",
		})
		return
	}

	_, err = db.GetTask(task.ID)
	if err != nil {
		writeJSON(w, map[string]string{
			"error": "ошибка",
		})
		return
	}

	err = checkDate(&task)
	if err != nil {
		writeJSON(w, map[string]string{
			"error": err.Error(),
		})
		return
	}

	err = db.UpdateTask(&task)
	if err != nil {
		writeJSON(w, map[string]string{
			"error": err.Error(),
		})
		return
	}

	writeJSON(w, map[string]string{})
}