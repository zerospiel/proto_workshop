//go:generate mkdir -p ./third_party/google/type/ ./third_party/k8s.io/apimachinery/pkg/runtime/ ./generated/
//go:generate curl -sSL0 https://raw.githubusercontent.com/googleapis/googleapis/15c5e21948ff6fbe41f91bdf04f6252f91a12d59/google/type/decimal.proto -o ./third_party/google/type/decimal.proto
//go:generate curl -sSL0 https://raw.githubusercontent.com/kubernetes/apimachinery/2e9dd8bce465be00efa368aeea7eff14e6d87978/pkg/runtime/generated.proto -o ./third_party/k8s.io/apimachinery/pkg/runtime/generated.proto
//go:generate protoc --proto_path .:third_party --go_opt=paths=source_relative,Mk8s.io/apimachinery/pkg/runtime/generated.proto=k8s.io/apimachinery/pkg/runtime --go_out=generated example_imports.proto
package imports
