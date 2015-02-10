DEBUG_FLAG = $(if $(DEBUG),-debug)

deps:
	go get github.com/codegangsta/cli

run: deps
	go run ./*.go

install: deps
	go install
