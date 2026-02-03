package add

import (
	"cv-landing-cli/pkg/activity"

	"github.com/spf13/cobra"
)

func InitCmd(repo activity.ActivityRepository) *cobra.Command {
	cmd := cobra.Command{
		Use:   "add",
		Short: "Add a new item to DB",
	}
	cmd.AddCommand(ActivityCmd(repo))
	cmd.AddCommand(TagCmd())
	cmd.AddCommand(AttachmentCmd())
	return &cmd
}
