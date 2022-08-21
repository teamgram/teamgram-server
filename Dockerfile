FROM golang:1.17 AS builder
WORKDIR /app
COPY . .
RUN ./build.sh

FROM ubuntu:latest
WORKDIR /app
COPY --from=builder /app/teamgramd/ /app/
RUN chmod +x /app/docker/entrypoint.sh
RUN apt update -y && apt install -y ffmpeg && chmod +x /app/docker/entrypoint.sh
ENTRYPOINT /app/docker/entrypoint.sh
