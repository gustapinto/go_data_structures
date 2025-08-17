test:
	go test ./... -failfast -cover -count 1

test/verbose:
	go test ./... -failfast -cover -count 1 -v

fmt:
	go fmt ./...
