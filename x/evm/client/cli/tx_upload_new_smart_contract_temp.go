package cli

import (
    "strconv"
	
	"github.com/spf13/cobra"
    "github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/palomachain/paloma/x/evm/types"
)

var _ = strconv.Itoa(0)

func CmdUploadNewSmartContractTemp() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "upload-new-smart-contract-temp [abi] [bytecode] [constructor-input] [chain-id]",
		Short: "Broadcast message UploadNewSmartContractTemp",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
      		 argAbi := args[0]
             argBytecode := args[1]
             argConstructorInput := args[2]
             argChainID := args[3]
            
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUploadNewSmartContractTemp(
				clientCtx.GetFromAddress().String(),
				argAbi,
				argBytecode,
				argConstructorInput,
				argChainID,
				
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