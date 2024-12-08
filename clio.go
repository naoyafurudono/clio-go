package clio


import (
	"context"
	"encoding/json"
	"fmt"

	"connectrpc.com/connect"
	"github.com/spf13/cobra"
)

// Generate a spf13/cobra command for a connect rpc.
// It will incorporate protovalidate or such kind of intercepter in the future development.
func RpcCommand[Req, Res any](
	ctx context.Context,
	rpc func(context.Context, *connect.Request[Req]) (*connect.Response[Res], error),
	use, short, long string,
	reqData *string) *cobra.Command {
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
				return err
			}
			out, err := json.Marshal(res.Msg)
			if err != nil {
				return err
			}
			fmt.Println(string(out))
			return nil
		},
	}
}
