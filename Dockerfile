FROM golang:1.12 as builder
WORKDIR /module
COPY src /module
RUN CGO_ENABLED=0 GOOS=linux go build -o app main.go

FROM alpine:3  
WORKDIR /root/
COPY --from=builder /module/app /root
COPY public /root/public
CMD ["./app"] 