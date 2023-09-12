FROM golang:1.20-bullseye as builder
ARG APP
ADD . /app
WORKDIR /app/cmd/$APP
RUN go mod tidy
RUN go build -o app -ldflags '-s -w -extldflags "-static"' -tags "netgo fixtures" .

FROM gcr.io/distroless/base
ARG APP
WORKDIR /app
COPY --from=builder /app/cmd/$APP/app /app/app
COPY --from=builder /app/cmd/$APP/config.dist.yml /app/config.dist.yml
ENTRYPOINT ["/app/app"]