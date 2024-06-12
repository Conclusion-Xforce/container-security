FROM registry.access.redhat.com/ubi9/go-toolset:1.20 AS builder

COPY src/* .

RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /opt/app-root/server

FROM registry.access.redhat.com/ubi9/ubi-micro:latest
ARG SERVICE_PORT=8080

COPY --from=builder /opt/app-root/server /app/server

EXPOSE ${SERVICE_PORT}
ENV SERVICE_PORT=${SERVICE_PORT}

CMD ["/app/server"]