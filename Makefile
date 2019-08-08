SERVER_TAG=0.0.1
DB_TAG=0.0.1

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

suika-db: protoc
	@echo "Building suika-db"
	go build -o target/suika-db \
		github.com/archelangelo/grpc-istio-demo/src/suika-db

clean:
	go clean github.com/archelangelo/grpc-istio-demo/...
	cd target
	rm -f server client

docker-server: protoc dep
	@echo "Building server docker image"
	docker build -t archelangelo/grpc-demo-server:${SERVER_TAG} -f src/docker/server/Dockerfile src/.

docker-db: protoc dep
	@echo "Building suika-db docker image"
	docker build -t archelangelo/suika-db:${DB_TAG} -f src/docker/suika-db/Dockerfile src/.

.PHONY: client server protoc dep suika-db docker-server docker-db
