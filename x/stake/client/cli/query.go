package cli

// // NewPostsQueryCmd defines the command to query posts from a vendor_id
// func NewPostsQueryCmd() *cobra.Command {
// 	cmd := &cobra.Command{
// 		Use:   "posts [vendor-id]",
// 		Args:  cobra.ExactArgs(1),
// 		Short: "Query for posts by vendor ID",
// 		Long: strings.TrimSpace(
// 			fmt.Sprintf(`Query posts by vendor ID.
// Example:
// $ %s query curating posts 1
// `,
// 				version.AppName,
// 			),
// 		),
// 		RunE: func(cmd *cobra.Command, args []string) error {
// 			clientCtx := client.GetClientContextFromCmd(cmd)
// 			clientCtx, err := client.ReadQueryCommandFlags(clientCtx, cmd.Flags())
// 			if err != nil {
// 				return err
// 			}

// 			vendorID, err := strconv.ParseUint(args[0], 10, 32)
// 			if err != nil {
// 				return err
// 			}

// 			queryClient := types.NewQueryClient(clientCtx)
// 			req := &types.QueryPostsRequest{
// 				VendorId: uint32(vendorID),
// 			}

// 			res, err := queryClient.Posts(context.Background(), req)
// 			if err != nil {
// 				return err
// 			}

// 			return clientCtx.PrintOutput(res)
// 		},
// 	}

// 	flags.AddQueryFlagsToCmd(cmd)
// 	return cmd
// }

// // NewPostQueryCmd defines the command to query a post from a vendor_id,post_id
// func NewPostQueryCmd() *cobra.Command {
// 	cmd := &cobra.Command{
// 		Use:   "post [vendor-id] [post-id]",
// 		Args:  cobra.ExactArgs(2),
// 		Short: "Query for a post by vendor ID and post ID",
// 		Long: strings.TrimSpace(
// 			fmt.Sprintf(`Query post by vendor ID and post ID.
// Example:
// $ %s query curating post 1 123
// `,
// 				version.AppName,
// 			),
// 		),
// 		RunE: func(cmd *cobra.Command, args []string) error {
// 			clientCtx := client.GetClientContextFromCmd(cmd)
// 			clientCtx, err := client.ReadQueryCommandFlags(clientCtx, cmd.Flags())
// 			if err != nil {
// 				return err
// 			}

// 			vendorID, err := strconv.ParseUint(args[0], 10, 32)
// 			if err != nil {
// 				return err
// 			}
// 			postID := strings.TrimSpace(args[1])

// 			if postID == "" {
// 				return fmt.Errorf("invalid post id")
// 			}

// 			queryClient := types.NewQueryClient(clientCtx)
// 			req := &types.QueryPostRequest{
// 				VendorId: uint32(vendorID),
// 				PostId:   postID,
// 			}

// 			res, err := queryClient.Post(context.Background(), req)
// 			if err != nil {
// 				return err
// 			}

// 			return clientCtx.PrintOutput(res)
// 		},
// 	}

// 	flags.AddQueryFlagsToCmd(cmd)
// 	return cmd
// }
