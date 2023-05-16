bootstrap: main.go
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o bootstrap .

clean:
	rm -f bootstrap

.PHONY: bootstrap
