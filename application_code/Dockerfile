FROM golang:alpine AS builder
RUN mkdir /build
ADD main.go /build/
WORKDIR /build
RUN go mod init main
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -a -o main .

FROM alpine
COPY --from=builder /build/main .
EXPOSE 8080
CMD ["./main"]
