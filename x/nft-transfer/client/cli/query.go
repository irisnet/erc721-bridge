package cli

import (
	"encoding/json"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"

	"github.com/bianjieai/nft-transfer/client/cli"

	"github.com/irisnet/erc721-bridge/x/nft-transfer/types"
)

// GetQueryCmd returns the parent command for all erc20 CLI query commands
func GetQueryCmd(fn types.QueryTokenTrace) *cobra.Command {
	cmd := cli.GetQueryCmd()
	cmd.AddCommand(
		GetTokenMappingCmd(fn),
	)
	return cmd
}

// GetTokenMappingCmd queries a cross-chain token trace
func GetTokenMappingCmd(queryTokenTrace types.QueryTokenTrace) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "token-trace <class_id> <token_id>",
		Short: "Get a cross-chain token trace",
		Long:  "When the target chain uses the erc-721 port of nft-transfer to receive cross-chain nft, the system will deploy an erc-721 contract to map with the nft of the original chain. You can use this command to enter ibc_class_id and the nft_id of the original chain to query, and you will get a new nft generated across the chain",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			traceClassId, traceTokenId, err := queryTokenTrace(clientCtx, args[0], args[1])
			if err != nil {
				return err
			}

			bz, err := json.Marshal(map[string]string{
				"class_id": traceClassId,
				"token_id": traceTokenId,
			})
			if err != nil {
				return err
			}
			return clientCtx.PrintBytes(bz)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}
