package lamblocal_test

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"testing"

	"log/slog"

	"github.com/aws/aws-lambda-go/events"
	"github.com/fujiwara/lamblocal"
)

func TestRunCLI_OK(t *testing.T) {
	handler := func(ctx context.Context, payload events.CloudWatchEvent) (string, error) {
		lamblocal.Logger.Info("hello", slog.String("ID", payload.ID))
		return "OK", nil
	}
	src := strings.NewReader(`{"id":"c4d2be6d-7987-411c-ada8-2713c427a115"}`)
	if out, err := lamblocal.RunCLI(context.Background(), src, handler); err != nil {
		t.Fatal(err)
	} else {
		t.Log(out)
	}
}

func TestRunCLIError(t *testing.T) {
	handler := func(ctx context.Context, payload events.CloudWatchEvent) (string, error) {
		lamblocal.Logger.Info("hello", slog.String("ID", payload.ID))
		return "", fmt.Errorf("error event: %s", payload.ID)
	}
	src := strings.NewReader(`{"id":"dddccaaf-7d6e-4332-b072-f46e7ad4ee2b"}`)
	if _, err := lamblocal.RunCLI(context.Background(), src, handler); err == nil {
		t.Error("error expected")
	} else if !strings.HasPrefix(err.Error(), "error event:") {
		t.Error("unexpected error:", err)
	} else {
		lamblocal.Logger.Error(err.Error())
	}
}

func TestRunCLINoPayload(t *testing.T) {
	handler := func(ctx context.Context, _ interface{}) (string, error) {
		return "OK", nil
	}
	src := strings.NewReader(``)
	if out, err := lamblocal.RunCLI(context.Background(), src, handler); err != nil {
		t.Fatal(err)
	} else {
		t.Log(out)
	}
}

func TestRunWithPayload(t *testing.T) {
	result := "NG"
	handler := func(ctx context.Context, s string) (string, error) {
		result = s
		return "", nil
	}
	lamblocal.CLISrc = strings.NewReader(`"OK"`)
	lamblocal.Run(context.Background(), handler)
	if result != "OK" {
		t.Error("unexpected result:", result)
	}
}

func TestRunWithError(t *testing.T) {
	e := errors.New("error")
	handler := func(ctx context.Context, _ struct{}) (struct{}, error) {
		return struct{}{}, e
	}
	lamblocal.CLISrc = strings.NewReader(`{}`)
	err := lamblocal.RunWithError(context.Background(), handler)
	if err == nil {
		t.Error("error expected")
	}
	if err != e {
		t.Error("unexpected error:", err)
	}
}
