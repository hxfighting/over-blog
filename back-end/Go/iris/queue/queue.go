package queue

import (
	"blog/models"
	"blog/service"
	"time"
)

func HandleQueue() {
	for {
		select {
		case id := <-service.EmailChan:
			go models.HandleEmailQueue(id)
		default:
			time.Sleep(time.Second)
		}
	}
}
