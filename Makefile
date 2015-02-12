deps:
	go get github.com/codegangsta/cli

run:
	go run ./*.go

cp:
	go build ./*.go

install: deps
	go install
