FROM golang:1.21-alpine AS builder
RUN apk add --no-cache ca-certificates
WORKDIR /go/src/app
COPY . /go/src/app
RUN CGO_ENABLED=0 go build -o app main.go

FROM scratch
WORKDIR /go/src/app
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /go/src/app/app ./
COPY --from=builder /go/src/app/web/ ./web/
COPY --from=builder /go/src/app/app.env ./
ENV TZ=Asia/Jakarta
EXPOSE 8080
CMD ["./app"]