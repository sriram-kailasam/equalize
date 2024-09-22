.DEFAULT_GOAL := buildAndStart

buildAndStart:
	go build -o ./bin/equalize
	./bin/equalize
