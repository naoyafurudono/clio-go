package clio

import (
	"context"
	"encoding/json"
	"errors"
	"io"

	"connectrpc.com/connect"
	"github.com/spf13/cobra"
)

var (
	// RPCFailed is the error returned when the rpc returns error.
	RPCFailed = errors.New("clio: rpc failed")
	// CLIFailed is the error returned when the CLI failed to do some work.
	CLIFailed = errors.New("clio: cli failed")
)

// Generate a spf13/cobra command for a connect rpc.
// Generated command executes the rpc and it writes the response to io.Writer w.
// use, short, and long are command name, short explanation, and long explanation.
// It will incorporate protovalidate or such kind of intercepter in the future development.
func RpcCommand[Req, Res any](
	ctx context.Context,
	rpc func(context.Context, *connect.Request[Req]) (*connect.Response[Res], error),
	use, short, long string,
	reqData *string, w io.Writer) *cobra.Command {
	return &cobra.Command{
		Use:   use,
		Short: short,
		Long:  long,
		RunE: func(cmd *cobra.Command, args []string) error {
			var req Req
			json.Unmarshal([]byte(*reqData), &req)
			res, err := rpc(
				ctx,
				connect.NewRequest(&req),
			)
			if err != nil {
				return errors.Join(RPCFailed, err)
			}
			out, err := json.Marshal(res.Msg)
			if err != nil {
				return errors.Join(CLIFailed, err)
			}
			if _, err := w.Write(out); err != nil {
				return errors.Join(CLIFailed, err)
			}
			return nil
		},
	}
}
