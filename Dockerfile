# Base image
ARG BUILD_REPO
ARG BUILD_BASE_IMAGE
ARG BUILD_TAG=latest

ARG REPO
ARG BASE_IMAGE
ARG TAG=latest

FROM $BUILD_REPO/$BUILD_BASE_IMAGE:$BUILD_TAG as build

LABEL maintainer="lhmzhou"

# Install Git and Curl
RUN apt-get update && \
    apt-get install -y --force-yes \
      curl \
      git && \
    rm -rf /var/lib/apt/lists/*

# Install Go
RUN curl https://storage.googleapis.com/golang/go1.8.3.linux-amd64.tar.gz | tar -C /usr/lib -xzvf-

# Source path of trans router code
RUN mkdir -p /go/src/purple-rain

# Set Go ENV Variables
ENV GOPATH=/go
ENV GOROOT=/usr/lib/go
ENV PATH=$PATH:$GOPATH

# Install Trans Router dependencies
RUN /usr/lib/go/bin/go get github.com/tools/godep
RUN /usr/lib/go/bin/go get golang.org/x/tools/cmd/godoc

# Add code
ADD . /go/src/purple-rain

# Install
RUN /usr/lib/go/bin/go install purple-rain

# Command
CMD /go/bin/purple-rain -v
