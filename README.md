# Experimental Microservices in Go

## Target features

- Universal API Gateway which support gRPC/REST/GraphQL client
- gPRC for service-service commnunication
- Multiregion support
- Easy to develop and deployment

## TODO

### Dev tools

- Research [earthly](https://github.com/earthly/earthly)
- Multi databases in one docker-compose service: https://github.com/mrts/docker-postgresql-multiple-databases/blob/master/Dockerfile
- Document generator
- Testing, mock service, Stress Test
- Lint, format, validator
- BloomRPC, Postman
- Versioning in monorepo (earthly)
- Advanced examples
- Goenv: Gopath, goroot, gobin,.. -> docker
- Code generator:
    - dockerize

### Monitoring

- middleware: otel, metrics,...
- logging

### API Gateway

- Research grpc-proxy, gRPC-gateway, ory gateway
- 

## How to setup

- Install goenv and go 1.20.4
- Install docker, docker-compose
- Setup recommend VScode extentions for go
- Install earthly

`brew install earthly && earthly bootstrap`

- Install buf

`brew install bufbuild/buf/buf`
`buf --version`
`cd shared/proto`
`buf mod init`
`cd ../../`
`touch buf.gen.yaml` and add configuration
`buf generate shared/proto`

## Guide

- https://levelup.gitconnected.com/microservices-with-go-grpc-api-gateway-and-authentication-part-1-2-393ad9fc9d30
- https://blog.logrocket.com/guide-to-grpc-gateway/#using-grpc-gateway-with-gin