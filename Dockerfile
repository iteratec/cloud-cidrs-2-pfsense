ARG USER_ID_NOBODY=65534
FROM golang:1.14-buster AS builder
LABEL maintainer=nils.kuhn@iteratec.com

ENV GO111MODULE=on

WORKDIR /build

# Let's cache modules retrieval - those don't change so often
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

# install taskfile (https://taskfile.dev)
RUN curl -sL https://taskfile.dev/install.sh | sh
# build cloud-cidr-2-pfsense
RUN ./bin/task cloud-cidrs-2-pfsense-build

FROM alpine:3

COPY --chown=65534:65534 --from=builder /build/dist/cloud-cidrs-2-pfsense /app

USER $USER_ID_NOBODY

EXPOSE 8080
CMD ["/app"]