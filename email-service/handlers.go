package task

import (
	"context"
	"fmt"

	"github.com/hibiken/asynq"
)

func NewMux() *asynq.ServeMux {

	mux := asynq.NewServeMux()

	fmt.Println("registered send email task")
	mux.HandleFunc("SEND_EMAIL", sendEmail)

	return mux
}

func sendEmail(ctx context.Context, t *asynq.Task) error {
	fmt.Println("sending email")
	return nil
}
