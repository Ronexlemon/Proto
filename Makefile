gen:
	@protoc \
	 --proto_path=protobuf "protobuf/coffee_proto.proto" \
	 --go_out=services/genproto  --go_opt=paths=source_relative \
	 --go-grpc_out=services/genproto  --go-grpc_opt=paths=source_relative