From golang:1.23-alpine as builder

WORKDIR /go/src/app

COPY . .

RUN CGO_ENABLED=0 go build -o ingress-manager main.go

FROM alpine:latest

workdir /app

COPY --from=builder /go/src/app/ingress-manager .

CMD ["./ingress-manager"]