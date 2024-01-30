build:
	go fmt && go build -o ./lsmods cmd/main.go
run:
	./lsmods
tests:
	export KMODULE=/lib/modules/`uname -r`/kernel/net/tls/tls.ko && go test -v ./...
all:
	make build
	make tests
	make run
