package show

import "github.com/spf13/cobra"

func ActivityCmd() *cobra.Command {
	return &cobra.Command{
		Use: "activity",
	}
}
