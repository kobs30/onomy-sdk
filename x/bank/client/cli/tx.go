package cli

import (
	"github.com/spf13/cobra"

	"github.com/onomyprotocol/onomy-sdk/client"
	"github.com/onomyprotocol/onomy-sdk/client/flags"
	"github.com/onomyprotocol/onomy-sdk/client/tx"
	sdk "github.com/onomyprotocol/onomy-sdk/types"
	"github.com/onomyprotocol/onomy-sdk/types/msgservice"
	"github.com/onomyprotocol/onomy-sdk/x/bank/types"
)

// NewTxCmd returns a root CLI command handler for all x/bank transaction commands.
func NewTxCmd() *cobra.Command {
	txCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Bank transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	txCmd.AddCommand(NewSendTxCmd())

	return txCmd
}

// NewSendTxCmd returns a CLI command handler for creating a MsgSend transaction.
func NewSendTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "send [from_key_or_address] [to_address] [amount]",
		Short: `Send funds from one account to another. Note, the'--from' flag is
ignored as it is implied from [from_key_or_address].`,
		Args: cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.Flags().Set(flags.FlagFrom, args[0])
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			toAddr, err := sdk.AccAddressFromBech32(args[1])
			if err != nil {
				return err
			}

			coins, err := sdk.ParseCoinsNormalized(args[2])
			if err != nil {
				return err
			}

			msg := types.NewMsgSend(clientCtx.GetFromAddress(), toAddr, coins)
			svcMsgClientConn := &msgservice.ServiceMsgClientConn{}
			msgClient := types.NewMsgClient(svcMsgClientConn)
			_, err = msgClient.Send(cmd.Context(), msg)
			if err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), svcMsgClientConn.GetMsgs()...)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
