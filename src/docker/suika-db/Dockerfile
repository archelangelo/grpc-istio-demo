FROM golang:alpine AS builder

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git

ENV PKG_NAME=github.com/archelangelo/grpc-istio-demo/src
ENV PKG_PATH=${GOPATH}/src/${PKG_NAME}

WORKDIR ${PKG_PATH}
COPY . .

RUN CGO_ENABLED=0 go build -o /go/bin/suika-db suika-db/main.go

FROM scratch
# Copy our static executable.
COPY --from=builder /go/bin/suika-db /go/bin/suika-db
EXPOSE 50051
ENV SUIKA_DB_PORT=50051
# Run the hello binary.
ENTRYPOINT ["/go/bin/suika-db"]