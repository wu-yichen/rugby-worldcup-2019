FROM yichen1987/ci as builder
LABEL maintainer="Yichen Wu <wuyichen1987@gmail.com>"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY ./ .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o wcwhen .

FROM alpine:latest as certs

RUN apk --update add ca-certificates

FROM scratch

WORKDIR /root/

COPY --from=builder /app/wcwhen .
COPY --from=builder /app/.env .
COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

ENTRYPOINT ["./wcwhen"]
