//go:generate rm -fr ./helloworldpb
//go:generate mkdir -p ./helloworldpb
//go:generate protoc --proto_path ../01_intro/ --go_opt=paths=source_relative --go_out=helloworldpb --go-grpc_opt=paths=source_relative --go-grpc_out=helloworldpb --twirp_opt=paths=source_relative --twirp_out=helloworldpb --go-vtproto_out=features=marshal+unmarshal+size,paths=source_relative:helloworldpb ../01_intro/helloworld/helloworld.proto
package plugins
