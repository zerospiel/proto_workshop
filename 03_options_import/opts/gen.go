//go:generate protoc --proto_path . --go_opt=module=github.com/zerospiel/proto_workshop/options/opts --go_out=. example_opts.proto
//go:generate mkdir -p generated_relative
//go:generate protoc --proto_path . --go_opt=paths=source_relative --go_out=generated_relative example_opts.proto
//go:generate mkdir -p generated_m_opt
//go:generate protoc --proto_path . --go_opt=paths=source_relative,Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types --go_out=generated_m_opt example_opts.proto
package opts
