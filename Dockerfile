FROM golang:1.15


    
WORKDIR /app

ENV PROTOC_ZIP=protoc-3.13.0-linux-x86_64.zip

RUN apt-get update && apt-get install -y unzip
RUN curl -OL https://github.com/protocolbuffers/protobuf/releases/download/v3.13.0/$PROTOC_ZIP \
    && unzip -o $PROTOC_ZIP -d /usr/local bin/protoc \
    && unzip -o $PROTOC_ZIP -d /usr/local 'include/*' \ 
    && rm -f $PROTOC_ZIP


COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . ./

RUN go build 

ENTRYPOINT go run server.go

#EXPOSE 7000

#ENTRYPOINT ["/go-roava"]
#CMD ["./server.go"]
