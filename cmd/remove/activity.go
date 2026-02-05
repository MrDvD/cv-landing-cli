package remove

import (
	"cv-landing-cli/pkg/activity"

	"github.com/spf13/cobra"
)

func fillActivityId(cmd cobra.Command) (cobra.Command, *int) {
	activityId := new(int)
	cmd.Flags().IntVar(activityId, "id", 0, "id of activity")
	cmd.MarkFlagRequired("id")
	return cmd, activityId
}

func ActivityCmd(repo activity.ActivityRepository) *cobra.Command {
	cmd := cobra.Command{
		Use: "activity",
	}
	cmd, activityId := fillActivityId(cmd)
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		return repo.Remove(*activityId)
	}
	return &cmd
}
