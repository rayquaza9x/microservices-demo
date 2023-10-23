FROM golang:1.20.4

WORKDIR /app

COPY ./services/product/go.mod .
# COPY ./services/product/go.sum .

# COPY ./cmd/product /app/cmd/product
COPY ./services/product /app/services/product
COPY ./shared /app/shared

WORKDIR /app/services/product
# CMD ["go", "run", "cmd/product/main.go"]
# CMD ["go", "run", "cmd/main.go"]

COPY ./containers/product/run.sh .
CMD ["bash", "run.sh"]