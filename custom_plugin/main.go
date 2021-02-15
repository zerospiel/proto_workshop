package main

import (
	"flag"
	"fmt"
	"os"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/pluginpb"
)

var (
	helloF = flag.String("hello", "default_value", "output hello")
	skipF  = flag.Bool("dryrun", false, "skip actual generating")
)

func main() {
	flag.Parse()

	protogen.Options{
		ParamFunc: flag.CommandLine.Set,
	}.Run(func(p *protogen.Plugin) error {
		p.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)

		for _, f := range p.Files {
			if f.Proto != nil {
				fmt.Fprintf(os.Stderr, "protoc-gen-mfp: processing %s\n", *f.Proto.Name)
			}
			if !f.Generate {
				continue
			}

			g := p.NewGeneratedFile(string(f.GoPackageName+".pb.mfp.go"), f.GoImportPath)
			g.P("// Code generated by protoc-gen-mfp. DO NOT EDIT.")
			g.P()
			g.P("package ", f.GoPackageName)
			g.P()

			g.P("func MyHelloWorld() {")
			g.P(
				protogen.GoImportPath("fmt").Ident("Printf"),
				`("(%s!\n)", "`+*helloF+`")`,
			)
			g.P("}")
			g.P()

			if *skipF {
				bb, _ := g.Content()
				fmt.Fprintf(os.Stderr, "Will generate:\n%s\n", string(bb))
				g.Skip()
			}
		}

		return nil
	})
}
