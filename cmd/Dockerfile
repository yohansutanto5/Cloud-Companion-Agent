FROM alpine:latest
WORKDIR /app
COPY . .
ENV config="config"
ENV GIN_MODE="test"
CMD ["./cloud-companion-agent"]
# CMD ["sh", "-c", "while :; do sleep 2073600; done"]