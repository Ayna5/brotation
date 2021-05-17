BIN := "./bin/bannersrotation"
PROTOPATH:=$(GOPATH)/src/github.com/Ayna5/bannersRotation/api/banners-rotation

.PHONY:  generate migrate lint install-lint-deps test build run up compose down

# set if not set
ifeq ($(POSTGRES_DB),)
	POSTGRES_DB := user=test password=test dbname=brotation host=localhost port=6432 sslmode=disable
endif

generate:
	mkdir -p pkg/banners-rotation
	protoc --proto_path=$(GOPATH)/src --go_out=. --go-grpc_out=. $(PROTOPATH)/banners_rotation.proto

migrate:
	goose -dir "./migrations" postgres "postgres://test:test@localhost:6432/brotation?sslmode=disable" up

lint:
	golangci-lint run ./...

install-lint-deps:
	(which golangci-lint > /dev/null) || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin v1.33.0

lint: install-lint-deps
	golangci-lint run ./...

test:
	go test -race -count 100 -timeout=30s ./internal/...

build:
	go build -v -o $(BIN) -ldflags "$(LDFLAGS)" ./cmd

run: build
	$(BIN) -config ./configs/config.toml

up:
	docker-compose up -d --build

compose:
	docker-compose -f docker-compose.yml up --build -d
	docker-compose ps -a

down:
	docker-compose down