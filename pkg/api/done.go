package api

import (
	"net/http"
	"time"

	"main/pkg/db"
)

func taskDoneHandler(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")

	if id == "" {
		writeJSON(w, map[string]string{
			"error": "ошибка",
		})
		return
	}

	task, err := db.GetTask(id)
	if err != nil {
		writeJSON(w, map[string]string{
			"error": err.Error(),
		})
		return
	}

	if task.Repeat == "" {
		err = db.DeleteTask(id)
	} else {
		next, err := NextDate(
			time.Now(),
			task.Date,
			task.Repeat,
		)
		if err != nil {
			writeJSON(w, map[string]string{
				"error": err.Error(),
			})
			return
		}

		err = db.UpdateDate(next, id)
	}

	if err != nil {
		writeJSON(w, map[string]string{
			"error": err.Error(),
		})
		return
	}

	writeJSON(w, map[string]string{})
}