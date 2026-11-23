PLATFORM=$(shell uname -s | tr '[:upper:]' '[:lower:]')
ifndef VERSION
	VERSION := $(shell git describe --tags --abbrev=0 2>/dev/null || echo v0.0.0)
endif
LDFLAGS := -X github.com/moov-io/x12.Version=$(VERSION)

USERID := $(shell id -u $$USER)
GROUPID:= $(shell id -g $$USER)

.PHONY: check
check:
ifeq ($(OS),Windows_NT)
	@echo "Skipping checks on Windows, currently unsupported."
else
	@wget -O lint-project.sh https://raw.githubusercontent.com/moov-io/infra/master/go/lint-project.sh
	@chmod +x ./lint-project.sh
	COVER_THRESHOLD=80.0 ./lint-project.sh
	@rm -rf cmd/x12/output
endif

dist: clean
ifeq ($(OS),Windows_NT)
	CGO_ENABLED=1 GOOS=windows go build -ldflags "$(LDFLAGS)" -o bin/x12.exe github.com/moov-io/x12/cmd/x12
else
	CGO_ENABLED=1 GOOS=$(PLATFORM) go build -ldflags "$(LDFLAGS)" -o bin/x12-$(PLATFORM)-amd64 github.com/moov-io/x12/cmd/x12
endif

.PHONY: clean
clean:
ifeq ($(OS),Windows_NT)
	@echo "Skipping cleanup on Windows, currently unsupported."
else
	@rm -rf cover.out coverage.txt misspell* staticcheck*
	@rm -rf ./bin/ openapi-generator-cli-*.jar x12.db ./storage/ lint-project.sh
	@rm -rf cmd/x12/output
endif

.PHONY: cover-test cover-web
cover-test:
	go test -coverprofile=cover.out ./...
cover-web:
	go tool cover -html=cover.out
