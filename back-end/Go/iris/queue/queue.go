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
		case id := <-service.ReplyEmailChan:
			go models.HandleReplyContactEmail(id)
		default:
			time.Sleep(time.Second)
		}
	}
}
