package main

import (
	"github.com/hibiken/asynq"
	task "github.com/sahilsk11/asynq/refresh-service"
)

func main() {
	redisAddr := "localhost:6379"
	r := asynq.RedisClientOpt{Addr: redisAddr}

	mux := task.NewMux()

	srv := asynq.NewServer(r, asynq.Config{
		Concurrency: 1,
	})

	srv.Run(mux)
}
