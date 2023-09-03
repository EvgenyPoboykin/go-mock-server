run:
	go run .

build:
	go build

d-build:
	docker build -t mock-server-app .

d-run:
	docker run -p 3000:3000 -it --rm --name mock-server-running-app mock-server-app
