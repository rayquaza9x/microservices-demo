FROM golang:1.20.4

WORKDIR /app

COPY ./services/api_gateway/go.mod .
COPY ./services/api_gateway/go.sum .

# COPY ./cmd/api_gateway /app/cmd/api_gateway
COPY ./services/api_gateway /app/services/api_gateway
COPY ./shared /app/shared

# CMD ["go", "run", "cmd/api_gateway/main.go"]
WORKDIR /app/services/api_gateway
# CMD ["go", "run", "cmd/main.go"]

COPY ./containers/api_gateway/run.sh .
CMD ["bash", "run.sh"]
