package api

import (
	"log"
	"net/http"

	"main/pkg/db"
)

type TaskResp struct {
	Tasks []*db.Task `json:"tasks"`
}

func tasksHandler(res http.ResponseWriter, req *http.Request) {
	tasks, err := db.Tasks(50)
	if err != nil {
		log.Println(err)

		writeJSON(res, map[string]string{
			"error": "не удалось получить список задач",
		}, http.StatusInternalServerError)
		return
	}

	writeJSON(res, TaskResp{
		Tasks: tasks,
	}, http.StatusOK)
}

func getTaskHandler(res http.ResponseWriter, req *http.Request) {
	id := req.FormValue("id")

	if id == "" {
		writeJSON(res, map[string]string{
			"error": "не указан id задачи",
		}, http.StatusBadRequest)
		return
	}

	task, err := db.GetTask(id)
	if err != nil {
		log.Println(err)

		writeJSON(res, map[string]string{
			"error": "задача не найдена",
		}, http.StatusNotFound)
		return
	}

	writeJSON(res, task, http.StatusOK)
}