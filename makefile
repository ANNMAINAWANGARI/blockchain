test:
	go test ./...

build:
	go build -o ./bin/customBlockchain

run: build
	./bin/customBlockchain