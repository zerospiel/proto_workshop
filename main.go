package main

import (
	"encoding/json"
	"fmt"

	pbws "github.com/zerospiel/proto_workshop/gen"
	"google.golang.org/protobuf/encoding/protojson"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

type kek struct {
	V int64 `json:"v"`
}

func main() {
	msg := &pbws.SomeMsg{Iv: &wrapperspb.Int32Value{}}
	msg.Iv.Value = 10000
	bb, err := protojson.Marshal(msg)
	if err != nil {
		panic(err)
	}
	var rj json.RawMessage
	err = json.Unmarshal(bb, &rj)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(rj))

	kv := &kek{V: 10000}
	bb, err = json.Marshal(&kv)
	if err != nil {
		panic(err)
	}
	var rj2 json.RawMessage
	err = json.Unmarshal(bb, &rj2)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(rj2))
}
