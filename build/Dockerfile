FROM golang:latest as builder

WORKDIR /app

RUN export GO111MODULE=on

COPY ./go.mod ./
COPY ./go.sum ./
RUN go mod download 
COPY . .

RUN CGO_ENABLED=0 go build -v -o ./url_shortener ./cmd/url_shortener

FROM alpine:latest

COPY --from=builder /app/url_shortener .
COPY --from=builder /app/config.yaml ./config.yaml

EXPOSE 8080 7080 7081 9000

ENTRYPOINT ["./url_shortener"]