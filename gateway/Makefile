# ==================================================================================== #
# HELPERS
# ==================================================================================== #
GIT_COMMIT := $(shell git rev-list -1 HEAD)
## make help OR make: Show the possible commands within the make file
.PHONY: help
help:
	@echo 'List of available commands:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

## make tidy: format code and tidy modfile
.PHONY: tidy
tidy:
	go fmt ./...
	go mod tidy -v
	go vet ./...
	go run honnef.co/go/tools/cmd/staticcheck@latest -checks=all,-ST1000,-U1000 ./...
	go test -race -vet=off ./...
	go mod verify

## make run: Serve the application without linting or formatting
.PHONY: run-debug
run-debug:
	@echo $(GIT_COMMIT)
	@export GIN_MODE=debug && go build -ldflags="-X 'main.Version=$(GIT_COMMIT)'" -o multiservice_gateway && ./multiservice_gateway

## make release: Serve the application without linting or formatting
.PHONY: run-release
run-release:
	@echo $(GIT_COMMIT)
	@export GIN_MODE=release && go release -ldflags="-X 'main.Version=$(GIT_COMMIT)'" -o multiservice_gateway && ./multiservice_gateway

## doc: Serve the documentation at :6060
.phony: doc
doc:
	@go mod tidy
	@go install golang.org/x/tools/cmd/godoc@latest
	godoc -http=:6060

## swag: Serve the swagger documentation
.phony: swag
swag:
	@go install github.com/swaggo/swag/cmd/swag@latest
	@swag init
	@go run main.go