run-docker:
	docker-compose up

# bin running
run:
	./bin/checkpass

# build binaries
build:
	GOOS=linux GOARCH=amd64 go build -trimpath -ldflags "-s" -o ./bin/checkpass *.go

build-win:
	GOOS=windows go build -trimpath -ldflags "-s" -o ./bin/checkpass.exe *.go

build-docker:
	docker-compose up --build
