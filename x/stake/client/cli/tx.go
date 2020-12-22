package cli

// // NewPostTxCmd returns the post command
// func NewPostTxCmd() *cobra.Command {

// 	cmd := &cobra.Command{
// 		Use:   "post [vendor-id] [post-id] [body] [reward_address] --from [key]",
// 		Args:  cobra.MinimumNArgs(3),
// 		Short: "Register a post",
// 		Long: strings.TrimSpace(
// 			fmt.Sprintf(`Create a post.
// Example:
// $ %s tx curating post 1 "2" "body" --from mykey
// `,
// 				version.AppName,
// 			),
// 		),
// 		RunE: func(cmd *cobra.Command, args []string) error {
// 			clientCtx := client.GetClientContextFromCmd(cmd)
// 			clientCtx, err := client.ReadTxCommandFlags(clientCtx, cmd.Flags())
// 			if err != nil {
// 				return err
// 			}

// 			creator := clientCtx.GetFromAddress()

// 			vendorID, err := strconv.ParseUint(args[0], 10, 32)
// 			if err != nil {
// 				return err
// 			}

// 			postID := args[1]
// 			body := args[2]

// 			var rewardAddrStr string
// 			var rewardAddr sdk.AccAddress
// 			if len(args) > 3 {
// 				rewardAddrStr = args[3]
// 			}
// 			if rewardAddrStr != "" {
// 				rewardAddr, err = sdk.AccAddressFromBech32(rewardAddrStr)
// 				if err != nil {
// 					return err
// 				}
// 			}
// 			msg := types.NewMsgPost(uint32(vendorID), postID, creator, rewardAddr, body)
// 			if err := msg.ValidateBasic(); err != nil {
// 				return err
// 			}

// 			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
// 		},
// 	}
// 	flags.AddTxFlagsToCmd(cmd)
// 	return cmd
// }