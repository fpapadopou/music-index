help:
	@echo "Usage: 'make <target>' where <target> is one of:"
	@echo "  ci       to run CI tests."
	@echo "  coverage to generate coverage report."
	@echo "  lint     to check code linting."

ci:
	go test ./... -race

coverage:
	go test -coverprofile=coverage.txt ./...

lint:
	golint ./...