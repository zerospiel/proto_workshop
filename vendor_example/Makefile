MODULE:=$(dir $(shell go list -m))

buf:
	@buf generate

protoc:
	protoc \
		--proto_path vendored_proto \
		--go_out=module=$(MODULE:/=):. \
		--go-grpc_out=module=$(MODULE:/=):. \
		$(shell find . -name "*.proto")
