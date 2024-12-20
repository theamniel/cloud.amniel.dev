# General info
GOOS 						:= $(shell go env GOOS)
OUTPUT_FOLDER 	:= ./build

# Config info
CONFIG_NAME   := config.toml
CONFIG_FILE 	:= ./config/$(CONFIG_NAME)
CONFIG_OUTPUT := $(OUTPUT_FOLDER)/$(CONFIG_NAME)

# Gateway info
GATEWAY_NAME 		:= cloud.gateway
GATEWAY_ENTRY 	:= ./gateway
GATEWAY_OUTPUT 	:= $(OUTPUT_FOLDER)/$(GATEWAY_NAME)

# Multimedia info
MULTIMEDIA_NAME 		:= cloud.multimedia
MULTIMEDIA_ENTRY 		:= ./services/multimedia
MULTIMEDIA_OUTPUT 	:= $(OUTPUT_FOLDER)/$(MULTIMEDIA_NAME)
MULTIMEDIA_PROTO	 	:= $(MULTIMEDIA_ENTRY)/protocols/*.proto

# Datastore info
DATASTORE_NAME 		:= cloud.datastore
DATASTORE_ENTRY 	:= ./services/datastore
DATASTORE_OUTPUT 	:= $(OUTPUT_FOLDER)/$(DATASTORE_NAME)
DATASTORE_PROTO 	:= $(DATASTORE_ENTRY)/protocols/*.proto

ifeq ($(GOOS), windows)
# Windows specific settings
	GATEWAY_OUTPUT 			:= $(GATEWAY_OUTPUT).exe
	MULTIMEDIA_OUTPUT 	:= $(MULTIMEDIA_OUTPUT).exe
	DATASTORE_OUTPUT 		:= $(DATASTORE_OUTPUT).exe
endif

generate-proto:
	protoc --go_out=. \
		--go_opt=paths=source_relative \
		--go-grpc_out=. \
		--go-grpc_opt=paths=source_relative \
		$(MULTIMEDIA_PROTO) $(DATASTORE_PROTO)

setup:
	@echo Updating dependency tree...
	go mod tidy
	go mod download
	@echo Updated dependency tree successfully.

format:
	@gofmt -w -s -l .

test:
	go test -covermode=atomic -parallel=4 -v -timeout=60s ./...

build-datastore:
	go build -o $(DATASTORE_OUTPUT) $(DATASTORE_ENTRY)

build-multimedia:
	go build -o $(MULTIMEDIA_OUTPUT) $(MULTIMEDIA_ENTRY)

build-gateway:
	go build -o $(GATEWAY_OUTPUT) $(GATEWAY_ENTRY)

build: build-datastore build-multimedia build-gateway