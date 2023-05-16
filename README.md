# lamblocal

[![Go Reference][1]][2]

[1]: https://pkg.go.dev/badge/github.com/fujiwara/lamblocal.svg
[2]: https://pkg.go.dev/github.com/fujiwara/lamblocal

## Description

lamblocal is a library that allows you to run AWS Lambda functions handler implemented by Go locally as a CLI command.

## Usage

```go
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
```

`lamblocal.Run()` executes the provided function, but the way it's executed depends on the current environment. Specifically, if the environment variable AWS_EXECUTION_ENV starts with "AWS_Lambda", or if AWS_LAMBDA_RUNTIME_API is set, the current environment is assumed to be AWS Lambda, and fn is executed as a Lambda function. This is achieved by calling lambda.Start(fn).

In all other environments, fn is executed as a CLI (Command Line Interface) function. In this case, the payload is read from the standard input (os.Stdin) and passed to fn. If the function returns an error, the error message is logged, and the program exits with an error code of 1.

### Logger

`lamblocal.Logger` is a logger that outputs to stderr as JSON format, using [slog](https://pkg.go.dev/golang.org/x/exp/slog).

## Limitation

Supports handler function interface `func (context.Context, interface{}) (interface{}, error)` only.

[aws-lambda-go](https://github.com/aws/aws-lambda-go) supports other handler function interface `func ()` and `func (ctx context.Context)`, and etc. but lamblocal does not support them.

## LICENSE

MIT
