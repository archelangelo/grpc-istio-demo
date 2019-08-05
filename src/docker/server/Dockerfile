FROM golang:alpine AS builder

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git

ENV PKG_NAME=github.com/archelangelo/grpc-istio-demo/src
ENV PKG_PATH=${GOPATH}/src/${PKG_NAME}

WORKDIR ${PKG_PATH}
COPY . .

RUN go build server/main.go -o /go/bin/server

FROM scratch
# Copy our static executable.
COPY --from=builder /go/bin/server /go/bin/server
EXPOSE 50051
# Run the hello binary.
ENTRYPOINT ["/go/bin/server"]