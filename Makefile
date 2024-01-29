build:
	go build -o ./lsmods cmd/main.go

run:
	./lsmods

test:
	go test -v ./...
