FROM golang:latest

WORKDIR /app
# ENV GOPATH=$GOPATH/app

COPY go.mod ./

COPY go.sum ./
RUN go mod tidy
RUN go mod vendor
RUN go get -d -v ./...
RUN go install -v ./...

COPY . ./

RUN go build -o /dist

EXPOSE 8080

CMD ["/dist"]
