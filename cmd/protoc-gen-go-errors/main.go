// Package main implements protoc-gen-go-errors plugin
// This is a backward-compatible wrapper around errgenkratos
//
// main 实现 protoc-gen-go-errors 插件
// 这是 errgenkratos 的向后兼容包装
package main

import (
	"flag"
	"fmt"
	"runtime/debug"

	"github.com/orzkratos/errgenkratos/erkgen"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/pluginpb"
)

var showVersion = flag.Bool("version", false, "print the version and exit")

func main() {
	flag.Parse()
	if *showVersion {
		if info, ok := debug.ReadBuildInfo(); ok {
			fmt.Printf("protoc-gen-go-errors %v (module: %s) (target: %s)\n", release, info.Main.Path, info.Path)
		} else {
			fmt.Printf("protoc-gen-go-errors %v\n", release)
		}
		return
	}
	var flags flag.FlagSet
	includeNested := flags.Bool("include_nested", false, "include nested enums")
	separateNamed := flags.Bool("separate_named", false, "use underscore_names")
	protogen.Options{
		ParamFunc: flags.Set,
	}.Run(func(gen *protogen.Plugin) error {
		gen.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)
		for _, f := range gen.Files {
			if !f.Generate {
				continue
			}
			erkgen.GenerateFile(gen, f, erkgen.Config{
				GeneratorName: "protoc-gen-go-errors (errgen)",
				IncludeNested: *includeNested,
				SeparateNamed: *separateNamed,
			})
		}
		return nil
	})
}
