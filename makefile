BINARY_NAME=gin
BINARY_NAME1=http
.DEFAULT_GOAL := run

build:
	CGO_ENABLED=0 GOOS=linux go build -o ./${BINARY_NAME}-server ./cmd/${BINARY_NAME}
	CGO_ENABLED=0 GOOS=linux go build -o ./${BINARY_NAME1}-server ./cmd/${BINARY_NAME1}

format:
	go fmt ...

linter: 
	golangci-lint run

images:
	docker build -f Dockerfile.${BINARY_NAME} --tag ${BINARY_NAME}---server .
	docker build -f Dockerfile.${BINARY_NAME1} --tag ${BINARY_NAME1}---server .

${BINARY_NAME}:
	docker run ${BINARY_NAME}---server

${BINARY_NAME1}:
	docker run ${BINARY_NAME1}--server

test:
	go test -v ./...