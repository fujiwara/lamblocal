package main

import (
	"context"

	"log/slog"

	"github.com/aws/aws-lambda-go/events"
	"github.com/fujiwara/lamblocal"
)

func myHandler(ctx context.Context, payload events.CloudWatchEvent) (string, error) {
	lamblocal.Logger.Info("hello", slog.String("ID", payload.ID))
	// do something
	return "OK", nil
}

func main() {
	lamblocal.Run(context.TODO(), myHandler)
}
