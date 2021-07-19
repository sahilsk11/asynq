package main

import (
	"fmt"
	"log"

	"github.com/hibiken/asynq"
)

func main() {
	redisAddr := "localhost:6379"
	r := asynq.RedisClientOpt{Addr: redisAddr}

	asynqClient := asynq.NewClient(r)
	defer asynqClient.Close()

	enqueueEmailTasks(asynqClient)
	enqueueRefreshTasks(asynqClient)
}

func enqueueEmailTasks(client *asynq.Client) {
	task := asynq.NewTask("SEND_EMAIL", nil)
	_, err := client.Enqueue(task, asynq.Queue("email"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("enqueued email task")
}

func enqueueRefreshTasks(client *asynq.Client) {
	task := asynq.NewTask("REFRESH", nil)
	_, err := client.Enqueue(task, asynq.Queue("refresh"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("enqueued refresh task")
}
