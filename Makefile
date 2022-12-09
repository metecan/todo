run:
	go run ./cmd/todo/todo.go

build:
	go build ./cmd/todo/todo.go

list:
	go run ./cmd/todo/todo.go -list

create:
	go run ./cmd/todo/todo.go -create Made By Makefile

complete:
	go run ./cmd/todo/todo.go -complete 1

remove:
	go run ./cmd/todo/todo.go -remove 1
