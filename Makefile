all: client server

dep:
	@echo "Install dependencies"
	cd src && glide install

protoc:
	@echo "Generating Go files"
	cd src/proto && protoc -I ${GOOGLEAPIS} -I . --go_out=plugins=grpc:. *.proto

server: protoc
	@echo "Building server"
	go build -o target/server \
		github.com/archelangelo/grpc-istio-demo/src/server

client: protoc
	@echo "Building client"
	go build -o target/client \
		github.com/archelangelo/grpc-istio-demo/src/client

clean:
	go clean github.com/archelangelo/grpc-istio-demo/...
	cd target
	rm -f server client

.PHONY: client server protoc dep
