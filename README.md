# Asynq Workers

This is an example repo to demonstrate Asynq's functionality.

## Running

First, start redis using `redis-server`.

To run the email worker, run:

```
go run cmd/refresh-worker/main.go
```

To run the refresh worker, run:

```
go run cmd/email-worker/main.go
```

To run the client, run:

```
go run cmd/client/main.go
```

The client adds one `SEND_EMAIL` job and one `REFRESH` job to the respective queues.