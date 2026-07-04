package api

import (
	"log"
	"net/http"
	"time"

	"main/pkg/db"
)

func taskDoneHandler(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")

	if id == "" {
		writeJSON(w, map[string]string{
			"error": "не указан id",
		}, http.StatusBadRequest)
		return
	}

	task, err := db.GetTask(id)
	if err != nil {
		log.Println(err)

		writeJSON(w, map[string]string{
			"error": err.Error(),
		}, http.StatusNotFound)
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
			log.Println(err)

			writeJSON(w, map[string]string{
				"error": err.Error(),
			}, http.StatusBadRequest)
			return
		}

		err = db.UpdateDate(next, id)
	}

	if err != nil {
		log.Println(err)

		writeJSON(w, map[string]string{
			"error": err.Error(),
		}, http.StatusInternalServerError)
		return
	}

	writeJSON(w, map[string]string{}, http.StatusOK)
}