package lamblocal

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	"log/slog"

	"github.com/aws/aws-lambda-go/lambda"
)

var Logger = slog.New(slog.NewJSONHandler(os.Stderr, nil))

var CLISrc io.Reader

func init() {
	CLISrc = os.Stdin // default
}

// Run runs a lambda handler func detect the environment (lambda or not) and run it.
func Run[T any, U any](ctx context.Context, fn func(context.Context, T) (U, error)) {
	err := RunWithError(ctx, fn)
	if err != nil {
		Logger.Error(err.Error())
		os.Exit(1)
	}
}

// RunWithError runs a lambda handler func detect the environment (lambda or not) and run it.
func RunWithError[T any, U any](ctx context.Context, fn func(context.Context, T) (U, error)) error {
	if strings.HasPrefix(os.Getenv("AWS_EXECUTION_ENV"), "AWS_Lambda") || os.Getenv("AWS_LAMBDA_RUNTIME_API") != "" {
		lambda.Start(fn)
	} else {
		out, err := RunCLI(ctx, CLISrc, fn)
		if err != nil {
			return err
		}
		json.NewEncoder(os.Stdout).Encode(out)
	}
	return nil
}

// RunCLI is a helper function for running a lambda handler func on CLI.
func RunCLI[T any, U any](ctx context.Context, src io.Reader, fn func(context.Context, T) (U, error)) (U, error) {
	payload := new(T)
	if err := json.NewDecoder(src).Decode(payload); err != nil {
		if err == io.EOF {
			return fn(ctx, *payload)
		}
		return *new(U), fmt.Errorf("failed to decode payload: %w", err)
	}
	return fn(ctx, *payload)
}
