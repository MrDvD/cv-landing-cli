package add

import (
	"cv-landing-cli/pkg/activity"

	"github.com/spf13/cobra"
)

func fillActivity(cmd cobra.Command) (cobra.Command, *activity.Activity) {
	activity := activity.Activity{}
	cmd.Flags().StringVar(&activity.Name, "name", "", "name for activity")
	cmd.MarkFlagRequired("name")
	var subtitle string
	cmd.Flags().StringVar(&subtitle, "subtitle", "", "subtitle for activity")
	if subtitle != "" {
		activity.Subtitle = &subtitle
	}
	cmd.Flags().StringVar(&activity.Description, "description", "", "description for activity")
	cmd.MarkFlagRequired("description")
	cmd.Flags().StringVar(&activity.Type, "type", "", "type of activity")
	cmd.MarkFlagRequired("type")
	var meta string
	cmd.Flags().StringVar(&meta, "meta", "", "meta label for activity")
	if meta != "" {
		activity.MetaLabel = &meta
	}
	cmd.Flags().StringVar(&activity.DateStart, "date-start", "", "start date of activity")
	cmd.MarkFlagRequired("date-start")
	var dateEnd string
	cmd.Flags().StringVar(&dateEnd, "date-end", "", "end date of activity")
	if dateEnd != "" {
		activity.DateEnd = &dateEnd
	}
	return cmd, &activity
}

func ActivityCmd(repo activity.ActivityRepository) *cobra.Command {
	cmd := cobra.Command{
		Use: "activity",
	}
	cmd, activity := fillActivity(cmd)
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		_, err := repo.Add(*activity)
		return err
	}
	return &cmd
}
