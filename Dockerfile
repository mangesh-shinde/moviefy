FROM golang:alpine3.18
WORKDIR /usr/src/app
COPY ./go.mod .
COPY ./go.sum .
RUN go mod download 
COPY . .
RUN go build -o main ./main.go
CMD ["./main"]