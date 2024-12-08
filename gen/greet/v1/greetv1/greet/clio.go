// Code generated by protoc-gen-clio-go. DO NOT EDIT.
//
// Source: greet/v1/greet.proto

package greetv1clio

import (
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