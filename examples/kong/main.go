package main

import (
	"context"

	"github.com/alecthomas/kong"
	"github.com/aws/aws-lambda-go/events"
	"github.com/fujiwara/lamblocal"
	"golang.org/x/exp/slog"
)

type CLI struct {
	Verbose bool   `help:"Verbose mode." default:"false" env:"VERBOSE"`
	Foo     string `help:"Foo." default:"foo" env:"FOO"`
}

func (c *CLI) Handler(ctx context.Context, payload events.CloudWatchEvent) (string, error) {
	lamblocal.Logger.Info("hello", slog.String("ID", payload.ID))
	// do something
	lamblocal.Logger.Info("foo", slog.String("foo", c.Foo))
	lamblocal.Logger.Info("verbose", slog.Bool("verbose", c.Verbose))
	return "OK", nil
}

func main() {
	var c CLI
	kong.Parse(&c)
	lamblocal.Run(context.TODO(), c.Handler)
}
