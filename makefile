gen:
	@protoc \
	--proto_path=protobuf "protobuf/game.proto" \
	--go_out=services/common/genproto --go_opt=paths=source_relative \
	--go-grpc_out=services/common/genproto --go-grpc_opt=paths=source_relative