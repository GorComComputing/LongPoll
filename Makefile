.PHONY: main
main: *.go deps
	go build -o main .
	./main


.PHONY:deps
deps:
#	go get github.com/gorilla/sessions



