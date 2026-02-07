FROM golang:1.21.13 AS builder
WORKDIR /app
COPY . .
RUN ./build.sh

FROM ubuntu:latest
RUN apt update -y && apt install -y ffmpeg psmisc && apt-get clean
WORKDIR /app
COPY --from=builder /app/teamgramd/ /app/
RUN chmod +x /app/docker/entrypoint.sh
ENTRYPOINT /app/docker/entrypoint.sh
