FROM golang:1.20.4

WORKDIR /app

COPY ./services/auth/go.mod .
COPY ./services/auth/go.sum .

# COPY ./cmd/auth /app/cmd/auth
COPY ./services/auth /app/services/auth
COPY ./shared /app/shared

WORKDIR /app/services/auth
# CMD ["go", "run", "cmd/auth/main.go"]
# CMD ["go", "run", "cmd/main.go"]

COPY ./containers/auth/run.sh .
CMD ["bash", "run.sh"]