package show

import (
	"cv-landing-cli/pkg/activity"

	"github.com/spf13/cobra"
)

func InitCmd(repo activity.ActivityRepository) *cobra.Command {
	cmd := cobra.Command{
		Use:   "show",
		Short: "Show items from DB",
	}
	cmd.AddCommand(ActivityCmd())
	cmd.AddCommand(TagCmd())
	cmd.AddCommand(AttachmentCmd())
	return &cmd
}
