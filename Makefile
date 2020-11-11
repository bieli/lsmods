build:
	go build -i main.go && mv main lsmods

run:
	./lsmods

test:
	go test -v ./...
