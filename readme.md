# Protoc basics workshop

<!-- TOC -->

- [Protoc basics workshop](#protoc-basics-workshop)
    - [Tooling](#tooling)
    - [Useful references](#useful-references)
    - [Procobuf and protoc dumb flows](#procobuf-and-protoc-dumb-flows)

<!-- /TOC -->

## Tooling

[Check out](./tools.md) versions that I used to be more consistent with.

## Useful references

1. [Protocol Buffers Version 3 Language Guide](<https://developers.google.com/protocol-buffers/docs/proto3>)
1. **[Protocol Buffers Version 3 Language Specification](<https://developers.google.com/protocol-buffers/docs/reference/proto3-spec>)**
1. [Protocol Buffers Version 3 Importing definitions](<https://developers.google.com/protocol-buffers/docs/proto3#importing_definitions>)
1. **[Style Guide](<https://developers.google.com/protocol-buffers/docs/style>)**
1. [Protobuf API Implementation Repository](<https://github.com/protocolbuffers/protobuf>)
1. [Protobuf API Golang Implementation v2 Repository](<https://github.com/protocolbuffers/protobuf-go>)
1. **[Generate Go Code](<https://developers.google.com/protocol-buffers/docs/reference/go-generated>)**
1. [gRPC Basics Tutorial](<https://grpc.io/docs/languages/go/basics/>)
1. [gRPC Server Reflection Protocol](<https://github.com/grpc/grpc/blob/master/doc/server-reflection.md>)
1. [gRPC-Go Server Reflection Symbols Problem](<https://github.com/grpc/grpc-go/issues/2328#issuecomment-424422612>)
1. [Lyft's tool for protoc plugins creation](<https://github.com/lyft/protoc-gen-star>)
1. [Buf Build Documentation](<https://docs.buf.build>)

## Procobuf and protoc dumb flows

For simplification purposes for checking out, I placed images of flow here:

Protobuf flow:

![protobuf_flow](./01_intro/pb.svg)

Protoc flow:

![protoc_flow](./02_plugins/protoc.svg)
