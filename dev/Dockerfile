##normally I'd use a multi step build
FROM golang:1.18-alpine

# move working directory
WORKDIR /app

# install packages
COPY go.mod ./
COPY go.sum ./
COPY ./ ./
RUN go mod download

RUN go build ./cmd/server
# start app
EXPOSE 8080
ENTRYPOINT ["./server"]
