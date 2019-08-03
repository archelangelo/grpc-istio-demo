protoc \
-I ../googleapis \
-I pingpong/ \
--go_out=plugins=grpc:pingpong \
ping-pong.proto
