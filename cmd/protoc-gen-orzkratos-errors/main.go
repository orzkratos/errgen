package main

import (
	"flag"
	"fmt"

	"github.com/orzkratos/errgen/internal/codegen"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/pluginpb"
)

var showVersion = flag.Bool("version", false, "print the version and exit")

func main() {
	flag.Parse()
	if *showVersion {
		fmt.Printf("protoc-gen-orzkratos-errors %v\n", release)
		return
	}
	var flags flag.FlagSet
	protogen.Options{
		ParamFunc: flags.Set,
	}.Run(func(gen *protogen.Plugin) error {
		gen.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)
		for _, f := range gen.Files {
			if !f.Generate {
				continue
			}
			codegen.GenerateFile(gen, f, codegen.Config{
				GeneratorName: "protoc-gen-orzkratos-errors (errgen)",
			})
		}
		return nil
	})
}
