package task

import (
	"context"
	"fmt"

	"github.com/hibiken/asynq"
)

func NewMux() *asynq.ServeMux {

	mux := asynq.NewServeMux()
	fmt.Println("registering refresh handler")
	mux.HandleFunc("REFRESH", refresh)

	return mux
}

func refresh(ctx context.Context, t *asynq.Task) error {
	fmt.Println("refreshing")
	return nil
}
