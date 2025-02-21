FROM debian:stable-slim

EXPOSE 80

WORKDIR /app
COPY ./bin/monitor-gateway /app/monitor-gateway

ENTRYPOINT ["/app/monitor-gateway"]