
build:
	go build ./cmd/server/main.go

debug:
	gdb ./main

run:
	go run ./cmd/server/main.go
