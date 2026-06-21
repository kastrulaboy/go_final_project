package api

import (
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
				"error": "ошибка",
			})
			return
		}

		err := db.DeleteTask(id)
		if err != nil {
			writeJSON(w, map[string]string{
				"error": err.Error(),
			})
			return
		}

		writeJSON(w, map[string]string{})

	default:
		http.Error(w, "ошибка", http.StatusMethodNotAllowed)
	}
}