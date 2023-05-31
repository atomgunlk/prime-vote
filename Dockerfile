FROM golang:1.20.1-alpine3.17 AS builder

WORKDIR /src/go
COPY go.mod go.sum ./

RUN go mod download

COPY ./cmd/ ./cmd
COPY ./internal/ ./internal

ENV APP_ENV=development
ENV APP_PORT=8000
ENV LOG_LEVEL=debug
ENV DB_HOST=127.0.0.1
ENV DB_PORT=5433
ENV DB_USERNAME=root
ENV DB_PASSWORD=root
ENV DB_NAME=primevote
ENV JWT_SECRET=testsecret1234

RUN go test ./cmd/prime-vote/handler -run "TestUnitHandler" -cover
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o /app "./cmd/prime-vote"

FROM alpine:3.17
RUN apk update && \
   apk add tzdata && \
   rm -rf /var/cache/apk/**

COPY --from=builder /app ./
ENTRYPOINT ["./app"]
