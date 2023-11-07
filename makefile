# ==================================================================================== #
# HELPERS
# ==================================================================================== #

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

# ==================================================================================== #
# DEVELOPMENT
# ==================================================================================== #

## tidy: tidy modfile
.PHONY: tidy
tidy:
	go fmt ./...
	go mod tidy -v
	go mod vendor

## upgrade: upgrade mofile
.PHONY: upgrade
upgrade:
	go get -u -v ./...
	go mod tidy
	go mod vendor

## audit: run quality control checks
.PHONY: audit
audit:
	go mod verify
	go vet ./...
	go run honnef.co/go/tools/cmd/staticcheck@latest -checks=all,-ST1000,-U1000 ./...
	go run golang.org/x/vuln/cmd/govulncheck@latest ./...
	go test -race -buildvcs -vet=off ./...

## test: run all tests
.PHONY: test
test:
	go test -v -race -buildvcs ./...

## test/cover: run all tests and display coverage
.PHONY: test/cover
test/cover:
	go test -v -race -buildvcs -coverprofile=/tmp/coverage.out ./...
	go tool cover -html=/tmp/coverage.out

# ==================================================================================== #
# BUILD
# ==================================================================================== #

## docker: build doker image
.PHONY: docker
docker:
	docker buildx build --platform=linux/amd64,linux/arm64 -t quay.io/phbpx/mockit:latest . --push

## mockit: build executable
.PHONY: mockit
mockit:
	go build -o mockit ./cmd/main.go