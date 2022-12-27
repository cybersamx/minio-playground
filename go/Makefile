PROJECT := mcgo
PROJECT_BIN := ./bin
SRC := .
GO_ENV := CGO_ENABLED=0

.PHONY: all
all: build

.PHONY: build
build:
	@-echo "Building $(PROJECT)."
	$(GO_ENV) go build -o $(PROJECT_BIN)/$(PROJECT) $(SRC)
