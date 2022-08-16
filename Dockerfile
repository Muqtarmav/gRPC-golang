FROM golang:1.18-alpine

WORKDIR /app


COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . ./

RUN go build 

#EXPOSE 7000

#ENTRYPOINT ["/go-roava"]
CMD ["go", "run", "./server.go"]
