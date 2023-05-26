package main

import (
	"assignment-3/routers"
	"assignment-3/utils"
	"net/http"
)

func main() {
	r := router.Router()

	go utils.CronJob()
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err)
	}
}
