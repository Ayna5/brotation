FROM golang:1.15.6

RUN mkdir -p /opt/integration-tests
WORKDIR /opt/integration-tests

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . /opt/integration-tests
CMD ["go", "test"]
