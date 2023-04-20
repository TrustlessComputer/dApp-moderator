FROM golang:1.18-bullseye as deps

RUN apt-get -y update && apt-get -y upgrade && \
    apt-get -y install git && \
    apt-get -y install make

ARG ENV=dev

ENV ENV=${ENV} \
    CGO_ENABLED=1

WORKDIR /app

COPY go.mod go.sum Makefile ./

RUN make init

RUN go mod download

FROM deps as builder
COPY  . .

RUN echo "✅ Build for Linux"; make build

# Distribution
FROM node:16-bullseye as nodejs-builder

WORKDIR /app
COPY ./tools /app/tools

RUN cd /app/tools/compile-contract/base-project \
    && npm install


FROM nodejs-builder as runner

WORKDIR /app
COPY --from=builder /app/tools /app/tools

COPY --from=builder /app/backend-api /app/backend-api
