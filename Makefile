all: install swagger run

install:
	go install -v

build:
	go build -o bin/main main.go

run:
	go run main.go

swagger:
	swag init -g main.go --output docs/habittracker