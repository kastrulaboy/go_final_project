package api

import (
	"log"
	"net/http"

	"main/pkg/db"
)

func taskHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case http.MethodPost:
		addTaskHandler(w, r)

	case http.MethodGet:
		getTaskHandler(w, r)

	case http.MethodPut:
		editTaskHandler(w, r)

	case http.MethodDelete:
		id := r.FormValue("id")

		if id == "" {
			writeJSON(w, map[string]string{
				"error": "не указан id задачи",
			}, http.StatusBadRequest)
			return
		}

		err := db.DeleteTask(id)
		if err != nil {
			log.Println(err)

			writeJSON(w, map[string]string{
				"error": err.Error(),
			}, http.StatusInternalServerError)
			return
		}

		writeJSON(w, map[string]string{}, http.StatusOK)

	default:
		http.Error(w, "метод не поддерживается", http.StatusMethodNotAllowed)
	}
}