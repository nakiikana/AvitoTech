server:
	go build -o server ./cmd/main.go
	./server

build:
	docker-compose build

run:
	docker-compose up

build and run: 
	docker-compose up --build

clean_server:
	rm -f server

stop:
	docker-compose down
