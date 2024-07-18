FROM golang:latest

WORKDIR /app
# ENV GOPATH=$GOPATH/app

COPY go.mod ./

COPY go.sum ./

ENV MONGO_URI=mongodb://localhost:27017
ENV GOROOT=/usr/local/go
ENV GOPATH=$PWD/bin
ENV GO_ENV=dev
ENV GIN_MODE=debug

RUN go mod tidy
RUN go mod vendor
RUN go get -d -v ./...
RUN go install -v ./...

COPY . ./

RUN go build -o /dist

EXPOSE 8080

CMD ["/dist"]
