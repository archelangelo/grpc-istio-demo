FROM golang:1.12

RUN curl https://glide.sh/get | sh

ENV PKG_NAME=github.com/archelangelo/grpc-istio-demo/src
ENV PKG_PATH=$GOPATH/src/$PKG_NAME
WORKDIR $PKG_PATH

COPY . $PKG_PATH/
RUN glide install

WORKDIR $PKG_PATH/server
RUN go install

EXPOSE 50051
CMD ["server"]