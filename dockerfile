FROM golang:latest
WORKDIR /dcard
COPY . .
RUN go mod download
RUN go build -o main .
EXPOSE 8081
CMD ["./main"]