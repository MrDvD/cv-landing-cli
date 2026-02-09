package show

import (
	"cv-landing-cli/pkg/activity"
	"cv-landing-cli/pkg/attachments"
	"cv-landing-cli/pkg/tags"

	"github.com/spf13/cobra"
)

func InitCmd(activityRepo activity.ActivityRepository, attachmentRepo attachments.AttachmentRepository, tagsRepo tags.TagsRepository) *cobra.Command {
	cmd := cobra.Command{
		Use:   "show",
		Short: "Show items from DB",
	}
	cmd.AddCommand(ActivityCmd(activityRepo))
	cmd.AddCommand(TagCmd(tagsRepo))
	cmd.AddCommand(AttachmentCmd(attachmentRepo))
	return &cmd
}
