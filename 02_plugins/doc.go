// Package plugins will be used for several purposes in our discussion:
//
// - what is and how to protoc-plugins (go vanilla, vtproto, twirp, etc)
//
// - flow of protoc plugins
//
// - how to write own protoc plugins in golang based on google.golang.org/protobuf/compiler/protogen
//
// - NB: do not use std out for anything other that https://github.com/protocolbuffers/protobuf/blob/master/src/google/protobuf/descriptor.proto//
//
package plugins
