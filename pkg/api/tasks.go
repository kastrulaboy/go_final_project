package api

import (
	"net/http"

	"main/pkg/db"
)

type TaskResp struct {
	Tasks []*db.Task `json:"tasks"`
}

func tasksHandler(res http.ResponseWriter, req *http.Request) {

	tasks, err := db.Tasks(50)
	if err != nil {
			writeJSON(res, map[string]string{
				"error": err.Error(),
			})
			return
	}

	writeJSON(res, TaskResp{
		Tasks: tasks,
	})
}

func getTaskHandler(res http.ResponseWriter, req *http.Request) {

	id := req.FormValue("id")

	if id == "" {
		writeJSON(res, map[string]string{
			"error": "ошибка",
		})
		return
	}

	task, err := db.GetTask(id)
	if err != nil {
		writeJSON(res, map[string]string{
			"error": err.Error(),
		})
		return
	}

	writeJSON(res, task)
}