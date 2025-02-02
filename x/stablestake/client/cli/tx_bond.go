package cli

import (
	"fmt"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/elys-network/elys/x/stablestake/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdBond() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "bond [amount]",
		Short: "Broadcast message bond",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			amount, ok := sdk.NewIntFromString(args[0])
			if !ok {
				return fmt.Errorf("unable to parse bonding amount")
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgBond(
				clientCtx.GetFromAddress().String(),
				amount,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
