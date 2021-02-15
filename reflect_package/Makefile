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
	# uncomment lines in protos before executing
	(rm -fr $(CURDIR)/pkg || true) && \
	mkdir -p $(CURDIR)/pkg && \
	protoc \
		--proto_path . \
		--go_out=paths=source_relative:pkg \
		--go-grpc_out=paths=source_relative:pkg \
		./foobar/foobar_svc.proto
	
fds:
	$(info file descriptor set of foobar.proto:)
	@protoc \
		--proto_path . \
		--descriptor_set_out /dev/stdout \
		./foobar/foobar.proto
	@echo

fds_opts:
	@protoc -I . -o ./tmp_wo.txt ./foobar/foobar.proto
	@protoc -I . -o ./tmp_with_imports.txt --include_imports ./foobar/foobar.proto
	@protoc -I . -o ./tmp_with_source_info.bin --include_source_info ./foobar/foobar.proto

clean:
	@rm ./tmp*