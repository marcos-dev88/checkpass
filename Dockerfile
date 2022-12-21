FROM golang:1.19

WORKDIR /verify-password-app/src/checkpass

COPY . .

ARG API_PORT

ENV GOPATH=/verify-password-app

ENV API_PORT ${API_PORT}

RUN GOOS=linux GOARCH=amd64 go build -trimpath -ldflags "-s" -o ./bin/checkpass *.go

ENTRYPOINT ["./bin/checkpass"]
