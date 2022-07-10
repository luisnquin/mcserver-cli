.PHONY: build

build:
	@go build -o ./build/mcserver .

run: 
	@./build/mcserver

clean:
	@rm -rf ./build/*

install:
	@go install .
