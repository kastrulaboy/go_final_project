package api

import (
	"fmt"
	"net/http"
)

func Init() {

	fmt.Print("Запускаю сервер")

	fs := http.FileServer(http.Dir("web"))

	http.HandleFunc("/api/task", taskHandler)
	http.HandleFunc("/api/nextdate", nextDayHandler)
	http.HandleFunc("/api/tasks", tasksHandler)
	http.HandleFunc("/api/task/done", taskDoneHandler)

	http.Handle("/", fs)

	err := http.ListenAndServe(":7540", nil)
	if err != nil {
		panic(err)
	}
}