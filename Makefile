.PHONY: build

build:
	GOOS=linux GOARCH=amd64 go build -o build/aws-penny-pincher
	cp config.yaml build/