TEST?=$$(go list ./... | grep -v 'vendor')
NAME=fusioncompute
BINARY=terraform-provider-${NAME}
VERSION=v0.1.1
OS_ARCH=darwin_amd64

default: build

build:
	go build -o ~/.terraform.d/plugins/darwin_amd64/${BINARY}

release:
	GOOS=linux GOARCH=amd64 go build -o ./bin/${BINARY}_${VERSION}_linux_amd64
	GOOS=linux GOARCH=arm64 go build -o ./bin/${BINARY}_${VERSION}_linux_arm64


