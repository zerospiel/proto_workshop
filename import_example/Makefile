MODULE:=$(shell go list -m)

generate:
	protoc \
		--proto_path . \
		--go_out=module=$(MODULE:/=):. \
		--go-grpc_out=module=$(MODULE:/=):. \
		$(shell find . -maxdepth 1 -name "*.proto")