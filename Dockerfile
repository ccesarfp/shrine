FROM golang:1.21.6-alpine3.19 as base
RUN apk update
WORKDIR /src/shrine
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -o shrine ./cmd/shrine

FROM alpine:3.14 as binary
WORKDIR /app
COPY --from=base /src/shrine/shrine /app
COPY --from=base /src/shrine/.env /app

CMD ["./shrine"]
