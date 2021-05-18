BIN := "./bin/bannersrotation"
PROTOPATH:=$(GOPATH)/src/github.com/Ayna5/bannersRotation/api/banners-rotation
LDFLAGS := -X main.release="develop" -X main.buildDate=$(shell date -u +%Y-%m-%dT%H:%M:%S) -X main.gitHash=$(GIT_HASH)

.PHONY:  generate migrate lint install-lint-deps test build run up compose down integration-test

# set if not set
ifeq ($(POSTGRES_DB),)
	POSTGRES_DB := user=test password=test dbname=brotation host=localhost port=6432 sslmode=disable
endif

generate:
	mkdir -p pkg/banners-rotation
	protoc --proto_path=$(GOPATH)/src --go_out=. --go-grpc_out=. $(PROTOPATH)/banners_rotation.proto

migrate:
	goose -dir "./migrations" postgres "postgres://test:test@db:5432/brotation?sslmode=disable" up

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

integration-test:
	docker-compose -f docker-compose.test.yml up --build -d ;\
	sleep 10 ;\
	docker-compose ps -a ;\
	test_status_code=0;\
	docker-compose -f docker-compose.test.yml run integration_tests go test ./integration-tests/... || test_status_code=$$?;\
	docker-compose -f docker-compose.test.yml down;\
	exit $$test_status_code;
