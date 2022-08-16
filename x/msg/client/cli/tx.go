package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/gotabit/gotabit/x/msg/types"
)

// NewTxCmd returns the transaction commands for the Msg module.
func NewTxCmd() *cobra.Command {
	txCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Msg transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	txCmd.AddCommand(
		GetCmdMsg(),
	)

	return txCmd
}

func GetCmdMsg() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "send [flags]",
		Long: "Send message to other user",
		Example: fmt.Sprintf(
			`$ %s tx msg send
				--to="gio1yx06xsqreefnhwmtu8ypd6vlatwxfqs9c2h2cq"
				--message="Example Message"`,
			version.AppName,
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			if len(args) != 2 {
				return sdkerrors.Wrapf(sdkerrors.Error{}, "invalid args length")
			}

			to := args[0]
			if err != nil {
				return err
			}
			message := args[1]
			if err != nil {
				return err
			}

			msg := types.NewMsgMsg(clientCtx.GetFromAddress().String(), to, message)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().AddFlagSet(FlagMsg())
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
