FROM ubuntu:22.04

ARG DEBIAN_FRONTEND="noninteractive"
RUN apt-get update -y

RUN apt install -y protobuf-compiler apt-utils asciidoc autoconf automake build-essential make
RUN apt install golang-goprotobuf-dev -y

WORKDIR /app/shared/proto

WORKDIR /app

COPY ./containers/protoc/Makefile .