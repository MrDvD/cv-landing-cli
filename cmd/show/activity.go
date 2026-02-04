package show

import (
	"cv-landing-cli/pkg/activity"
	"cv-landing-cli/pkg/view"
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

var tableConfig = view.TableConfig[activity.Activity]{
	Columns: []view.Column{
		{Header: "ID", Weight: 1},
		{Header: "NAME", Weight: 3},
		{Header: "SUBTITLE", Weight: 3},
		{Header: "DESCRIPTION", Weight: 6},
		{Header: "META", Weight: 2},
		{Header: "START DATE", Weight: 2},
		{Header: "END DATE", Weight: 2},
	},
	RowMapper: func(a activity.Activity) []string {
		return []string{
			fmt.Sprintf("%d", a.Id),
			a.Name,
			view.StrPtr(a.Subtitle),
			a.Description,
			view.StrPtr(a.MetaLabel),
			formatDate(a.DateStart),
			formatDate(view.StrPtr(a.DateEnd)),
		}
	},
}

func formatDate(date string) string {
	parsedDate, err := time.Parse(time.RFC3339, date)
	if err != nil {
		return date
	}
	return parsedDate.Format("Jan 2 2006")
}

func fillActivityType(cmd cobra.Command) (cobra.Command, *string) {
	var activityType string
	cmd.Flags().StringVar(&activityType, "type", "", "type of activities")
	cmd.MarkFlagRequired("type")
	return cmd, &activityType
}

func ActivityCmd(repo activity.ActivityRepository) *cobra.Command {
	cmd := cobra.Command{
		Use: "activity",
	}
	cmd, activityType := fillActivityType(cmd)
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		activities, err := repo.GetAllOfType(*activityType)
		if err != nil {
			return err
		}
		view.Table(activities, tableConfig)
		return nil
	}
	return &cmd
}
