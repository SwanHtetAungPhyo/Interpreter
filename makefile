SRC=./main.go
DEST=./bin/main

.PHONY: run build

run:
	@echo "Running the interpreter"
	go run $(SRC)

build:
	@echo "Building the application to binary file"
	go build -o $(DEST) $(SRC)