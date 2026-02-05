package remove

import "github.com/spf13/cobra"

func TagCmd() *cobra.Command {
	return &cobra.Command{
		Use: "tag",
	}
}
