FROM golang:1.20.14 AS builder
WORKDIR /app
COPY . .
RUN ./build.sh

FROM ubuntu:latest
WORKDIR /app
COPY --from=builder /app/teamgramd/ /app/
RUN apt update -y && apt install -y ffmpeg && apt-get clean && chmod +x /app/docker/entrypoint.sh
ENTRYPOINT /app/docker/entrypoint.sh
