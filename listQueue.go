package main

import (
	"net/http"
)

func ListQueueURLHandler(w http.ResponseWriter, r *http.Request) {
	result, _ := svc.ListQueues(nil)
	temp := ""
	for _, url := range result.QueueUrls {
		temp = temp + ", " + *url
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(temp))
}
