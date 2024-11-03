run:
	./dev.sh

start:
	./product.sh

build:
	go build -o app ./src/main.go 
	
run-build:
	export ENV_MODE=production
	./app

build-linux:
	GOOS=linux GOARCH=amd64 go build -o app ./src/main.go

build-docker:
	docker build -t server:latest .

run-docker:
	docker run --name server-web --network backend -p 8080:8080 -d server:latest
