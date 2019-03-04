protoc -I. -I%GoPath%/src -I build\extracted-include-protos\main\ --grpc-gateway_out=logtostderr=true:. ./src/main/proto/org/kai/example/helloworld.proto
protoc -I. -I%GoPath%/src -I build\extracted-include-protos\main\ --go_out=plugins=grpc:. ./src/main/proto/org/kai/example/helloworld.proto
