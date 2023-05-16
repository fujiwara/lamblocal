package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/fujiwara/lamblocal"
	"golang.org/x/exp/slog"
)

func myHandler(ctx context.Context, payload events.CloudWatchEvent) (string, error) {
	lamblocal.Logger.Info("hello", slog.String("ID", payload.ID))
	// do something
	return "OK", nil
}

func main() {
	lamblocal.Run(context.TODO(), myHandler)
}
