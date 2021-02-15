GOBIN?=$(GOPATH)/bin

install:
	go build -o $(GOBIN)/protoc-gen-mfp $(CURDIR)/main.go

generate:
	protoc -I ../reflect_package/foobar/ --mfp_out=hello="hello world":. ../reflect_package/foobar/foobar.proto

dryrun:
	protoc -I ../reflect_package/foobar/ --mfp_out=hello="hello world",dryrun=true:. ../reflect_package/foobar/foobar.proto