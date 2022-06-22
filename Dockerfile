FROM golang:latest

WORKDIR /app
# ENV GOPATH=$GOPATH/app

COPY go.mod ./

COPY go.sum ./

RUN go get -d -v ./...
RUN go install -v ./...

COPY . ./

RUN "ls"

RUN go build -o /dist

EXPOSE 8080

CMD ["/dist"]
