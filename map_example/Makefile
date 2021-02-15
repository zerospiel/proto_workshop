MODULE:=$(shell go list -m)

generate:
	$(eval map_opt="Mgithub.com/user/repo/api/foobar.proto=github.com/zerospiel/proto_workshop/map_example/external")
	protoc \
		--proto_path . \
		--go_out=$(map_opt),module=$(MODULE:/=):. \
		--go-grpc_out=$(map_opt),module=$(MODULE:/=):. \
		$(shell find . -name "*.proto")