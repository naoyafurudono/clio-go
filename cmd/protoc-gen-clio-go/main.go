package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"

	"google.golang.org/protobuf/compiler/protogen"
)

const (
	contextPackage = protogen.GoImportPath("context")

	usage   = `todo`
	version = "0.0.1"
)

func main() {
	if len(os.Args) == 2 && os.Args[1] == "--version" {
		fmt.Fprintln(os.Stdout, version)
		os.Exit(0)
	}
	if len(os.Args) == 2 && (os.Args[1] == "-h" || os.Args[1] == "--help") {
		fmt.Fprintln(os.Stdout, usage)
		os.Exit(0)
	}
	if len(os.Args) != 1 {
		fmt.Fprintln(os.Stderr, usage)
		os.Exit(1)
	}
	protogen.Options{}.Run(
		func(plugin *protogen.Plugin) error {
			for _, file := range plugin.Files {
				if file.Generate {
					generate(plugin, file)
				}
			}
			return nil
		},
	)
}

func generate(plugin *protogen.Plugin, f *protogen.File) {
	generatedFilenamePrefixToSlash := filepath.ToSlash(f.GeneratedFilenamePrefix)
	filepath := path.Join(
		path.Dir(generatedFilenamePrefixToSlash),
		// パッケージ専用のディレクトリを掘る
		string(f.GoPackageName),
		path.Base(generatedFilenamePrefixToSlash),
		"clio.go",
	)
	fmt.Fprint(os.Stderr, filepath)
	gf := plugin.NewGeneratedFile(filepath, f.GoImportPath)

	generatePreamble(gf, f)
	generateBody(gf, f)
}

func generatePreamble(g *protogen.GeneratedFile, file *protogen.File) {
	g.P("// Code generated by protoc-gen-clio-go. DO NOT EDIT.")
	g.P("//")
	if file.Proto.GetOptions().GetDeprecated() {
		g.P(file.Desc.Path(), " is a deprecated file.")
	} else {
		g.P("// Source: ", file.Desc.Path())
	}
	g.P()
	g.P("package ", file.GoPackageName+"clio")
	g.P()

}

func generateBody(g *protogen.GeneratedFile, _ *protogen.File) {
	g.P(`import (
			"context"

			"github.com/naoyafurudono/clio-go/gen/greet/v1/greetv1connect" // generated by protoc-gen-connect-go
			clio "github.com/naoyafurudono/clio-go"
			"github.com/spf13/cobra"
)

// CLI implementation (what we generate)
func NewGreetCommand(ctx context.Context, s greetv1connect.GreetServiceHandler) *cobra.Command {
			var greetService = &cobra.Command{
				Use:   "great",
				Short: "Important service.",
				Long:  "Important service.",
			}
			var reqData *string = greetService.PersistentFlags().StringP("data", "d", "{}", "request message represented as a JSON")

			greetServiceHello := clio.RpcCommand(ctx,
				s.Hello,
				"hello",
				"basic greeting",
				"basic greeting",
				reqData,
			)
			greetServiceThanks := clio.RpcCommand(ctx,
				s.Thanks,
				"thanks",
				"you cannot live alone",
				"you cannot live alone",
				reqData,
			)

			greetService.AddCommand(
				greetServiceHello,
				greetServiceThanks,
			)
			return greetService
}
		`)
}
