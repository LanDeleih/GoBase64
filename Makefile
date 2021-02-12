.PHONY: clean

DIR := bin

default: linux

clean: clean
	rm -rf ./$(DIR)

linux:
	GOOS=linux go build -o $(DIR)/goBase64 ./app/main.go

darwin:
	GOOS=darwin go build -o $(DIR)/goBase64 ./app/main.go

windows:
	GOOS=windows go build -o $(DIR)/goBase64.exe ./app/main.go

