FROM golang:1.22-alpine3.19 as build
RUN apk update
WORKDIR /src/shrine
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -o shrine ./cmd/shrine

FROM alpine:3.14 as app
WORKDIR /app
COPY --from=build /src/shrine/shrine /app
COPY --from=build /src/shrine/.env /app

RUN ["./shrine", "create:key"]

CMD ["./shrine", "up"]
