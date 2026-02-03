package show

import "github.com/spf13/cobra"

func AttachmentCmd() *cobra.Command {
	return &cobra.Command{
		Use: "attachment",
	}
}
