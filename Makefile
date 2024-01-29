build:
	go build -o ./lsmods cmd/main.go
run:
	./lsmods
tests:
	go test -v ./...
all:
	make build
	make tests
	make run
