build:
	go fmt && go build -o ./lsmods cmd/main.go
run:
	./lsmods
tests:
	export WORKSPACE=`pwd` && go test -v ./...
all:
	make build
	make tests
	make run
