package api

import (
	"encoding/json"
	"log"
	"net/http"

	"main/pkg/db"
)

func editTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task db.Task

	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		log.Println(err)
		writeJSON(w, map[string]string{
			"error": err.Error(),
		}, http.StatusBadRequest)
		return
	}

	if task.ID == "" {
		writeJSON(w, map[string]string{
			"error": "не указан id",
		}, http.StatusBadRequest)
		return
	}

	if task.Title == "" {
		writeJSON(w, map[string]string{
			"error": "не указан заголовок",
		}, http.StatusBadRequest)
		return
	}

	_, err = db.GetTask(task.ID)
	if err != nil {
		log.Println(err)
		writeJSON(w, map[string]string{
			"error": "задача не найдена",
		}, http.StatusNotFound)
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

	err = db.UpdateTask(&task)
	if err != nil {
		log.Println(err)
		writeJSON(w, map[string]string{
			"error": err.Error(),
		}, http.StatusInternalServerError)
		return
	}

	writeJSON(w, map[string]string{}, http.StatusOK)
}