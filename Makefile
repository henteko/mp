deps:
	go get github.com/codegangsta/cli
	go get howett.net/plist

install: deps
	go install
