protoc --go_out=.\src\message --proto_path=.\protobuild login.proto
protoc --go_out=plugins=grpc:.\src\message --proto_path=.\protobuild dbrpc.proto