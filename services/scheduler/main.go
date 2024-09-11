package main

import (
	"go-task-scheduler/pkg/natsutil"
	"log"
	"time"
)

func main() {
	nc := natsutil.ConnectNATS()
	defer nc.Close()

	ticker := time.NewTicker(10 * time.Second) // Schedule every 10 seconds

	for {
		select {
		case <-ticker.C:
			task := "Perform Task - " + time.Now().String()
			log.Println("Publishing task:", task)
			natsutil.PublishMessage(nc, "task.scheduled", task)
		}
	}
}
