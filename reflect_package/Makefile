MODULE:=$(dir $(shell go list -m))

protoc:
	(rm -fr $(CURDIR)/pkg || true) && \
	mkdir -p $(CURDIR)/pkg && \
	protoc \
		--proto_path . \
		--go_out=paths=source_relative:pkg \
		--go-grpc_out=paths=source_relative:pkg \
		$(shell find . -name "*.proto")

bad:
	(rm -fr $(CURDIR)/pkg || true) && \
	mkdir -p $(CURDIR)/pkg && \
	protoc \
		--proto_path . \
		--go_out=paths=source_relative:pkg \
		--go-grpc_out=paths=source_relative:pkg \
		./foobar/foobar_svc.proto