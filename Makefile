run: build start

test:
	@echo "Ganeca >> running test ..."
	@go test -v -race ./...

build:
	@echo "Ganeca >> building binaries ..."
	@go build -o bin/ganeca app.go
	@echo "Ganeca >> success"

start:
	@echo "Ganeca >> starting binaries ..."
	@./bin/ganeca