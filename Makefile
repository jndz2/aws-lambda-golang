.DEFAULT_GOAL := build

BINARY_NAME=lambdaHandler
BINARY_PATH=./bin
CODE_PATH=./lambda-function/main.go
INFRA_PATH=./lambda-infra

build:
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o "${BINARY_PATH}/${BINARY_NAME}" "${CODE_PATH}"

clean:
	go clean 
	rm -rf "${BINARY_PATH}/${BINARY_NAME}"

bootstrap:
	cd "${INFRA_PATH}" && cdk bootstrap

deploy:
	cd "${INFRA_PATH}" && cdk deploy 

destroy:
	cd "${INFRA_PATH}" && cdk destroy
