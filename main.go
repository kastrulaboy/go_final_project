package main

import (
	"main/pkg/db"
	"main/pkg/api"
//	nd "main/pkg/nextdate"
)

func main() {

	err := db.Init("scheduler.db")
		if err != nil {
			panic(err)
		}

	api.Init()

}