package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/irisnet/erc721-bridge/x/converter/types"
	"github.com/spf13/cobra"
)

// GetQueryCmd returns the parent command for all erc20 CLI query commands
func GetQueryCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the erc20 module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		GetClassPairsCmd(),
		GetClassPairCmd(),
		GetTokenTraceCmd(),
	)

	return cmd
}

// GetClassPairsCmd queries all registered token pairs
func GetClassPairsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "class-pairs",
		Short: "Gets registered class pairs",
		Long:  "Gets registered class pairs",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			clienCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clienCtx)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			req := &types.QueryClassPairsRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.ClassPairs(context.Background(), req)
			if err != nil {
				return err
			}

			return clienCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// GetClassPairCmd queries a registered token pair
func GetClassPairCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "class-pair CLASS",
		Short: "Get a registered class pair",
		Long:  "Get a registered class pair",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			req := &types.QueryClassPairRequest{
				Class: args[0],
			}

			res, err := queryClient.ClassPair(context.Background(), req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// GetTokenTraceCmd queries a cross-chain token trace
func GetTokenTraceCmd() *cobra.Command {
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

			queryClient := types.NewQueryClient(clientCtx)
			req := &types.QueryTokenTraceRequest{
				ClassId: args[0],
				TokenId: args[1],
			}

			data, err := queryClient.TokenTrace(cmd.Context(), req)
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(data)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}
