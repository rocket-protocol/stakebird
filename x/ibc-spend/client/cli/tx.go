package cli

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/version"

	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/public-awesome/stargaze/x/ibc-spend/types"
)

var (
	DefaultRelativePacketTimeoutTimestamp = uint64((time.Duration(10) * time.Minute).Nanoseconds())
)

const (
	flagPacketTimeoutTimestamp = "packet-timeout-timestamp"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	// this line is used by starport scaffolding # 1

	return cmd
}

// GetCmdSubmitProposal implements the command to submit an IBC community-pool-spend proposal
func GetCmdSubmitProposal() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "community-pool-ibc-spend [proposal-file]",
		Args:  cobra.ExactArgs(1),
		Short: "Submit a community pool IBC spend proposal",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Submit a community pool spend proposal along with an initial deposit.
The proposal details must be supplied via a JSON file.

Example:
$ %s tx gov submit-proposal community-pool-ibc-spend <path/to/proposal.json> --from=<key_or_address>

Where proposal.json contains:

{
  "title": "Community Pool IBC Spend",
  "description": "Pay me some tokens over IBC!",
  "recipient": "[prefix]1s5afhd6gxevu37mkqcvvsj8qeylhn0rz46zdlq",
  "amount": "1000stake",
  "sourceChannel": "channel-0",
  "timeout": 100,
  "deposit": "1000stake"
}
`,
				version.AppName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			proposal, err := ParseCommunityPoolIBCSpendProposalWithDeposit(clientCtx.JSONMarshaler, args[0])
			if err != nil {
				return err
			}

			amount, err := sdk.ParseCoinsNormalized(proposal.Amount)
			if err != nil {
				return err
			}

			deposit, err := sdk.ParseCoinsNormalized(proposal.Deposit)
			if err != nil {
				return err
			}

			from := clientCtx.GetFromAddress()

			content := types.NewCommunityPoolIBCSpendProposal(
				proposal.Title,
				proposal.Description,
				proposal.Recipient,
				amount,
				proposal.SourceChannel,
				proposal.Timeout)

			msg, err := govtypes.NewMsgSubmitProposal(content, deposit, from)
			if err != nil {
				return err
			}

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	return cmd
}

// ParseCommunityPoolIBCSpendProposalWithDeposit reads and parses a CommunityPoolSpendProposalWithDeposit from a file.
func ParseCommunityPoolIBCSpendProposalWithDeposit(cdc codec.JSONMarshaler, proposalFile string) (types.CommunityPoolIBCSpendProposalWithDeposit, error) {
	proposal := types.CommunityPoolIBCSpendProposalWithDeposit{}

	contents, err := ioutil.ReadFile(proposalFile)
	if err != nil {
		return proposal, err
	}

	if err = cdc.UnmarshalJSON(contents, &proposal); err != nil {
		return proposal, err
	}

	return proposal, nil
}