package edit

import (
	"cv-landing-cli/pkg/activity"

	"github.com/spf13/cobra"
)

func InitCmd(repo activity.ActivityRepository) *cobra.Command {
	cmd := cobra.Command{
		Use:   "edit",
		Short: "Edit an item from DB",
	}
	cmd.AddCommand(ActivityCmd(repo))
	cmd.AddCommand(TagCmd())
	cmd.AddCommand(AttachmentCmd())
	return &cmd
}
