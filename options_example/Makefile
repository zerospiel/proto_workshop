MODULE:=$(shell go list -m)

generate:
	protoc \
		--proto_path . \
		--go_out=module=$(MODULE:/=):. \
		--go-grpc_out=module=$(MODULE:/=):. \
		--java_out=. \
		--python_out=. \
		--csharp_out=. \
		$(shell find . -name "*.proto")