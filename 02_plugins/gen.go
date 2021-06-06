//go:generate mkdir -p helloworldpb
//go:generate protoc --proto_path ../ --go_opt=paths=source_relative --go_out=helloworldpb --go-grpc_opt=paths=source_relative --go-grpc_out=helloworldpb --twirp_opt=paths=source_relative --twirp_out=helloworldpb ../01_intro/helloworld/helloworld.proto
package plugins
