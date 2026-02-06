package edit

import (
	"cv-landing-cli/pkg/activity"
	"strconv"

	"github.com/spf13/cobra"
)

type flagVars struct {
	name        string
	subtitle    string
	description string
	a_type      string
	meta        string
	date_start  string
	date_end    string
}

func getPatches(cmd *cobra.Command, f *flagVars) []activity.EditField {
	var patches []activity.EditField

	if cmd.Flags().Changed("name") {
		patches = append(patches, activity.EditField{
			Name:  "name",
			Value: f.name,
		})
	}

	if cmd.Flags().Changed("subtitle") {
		patches = append(patches, activity.EditField{
			Name:  "subtitle",
			Value: f.subtitle,
		})
	}

	if cmd.Flags().Changed("description") {
		patches = append(patches, activity.EditField{
			Name:  "description",
			Value: f.description,
		})
	}

	if cmd.Flags().Changed("type") {
		patches = append(patches, activity.EditField{
			Name:  "type",
			Value: f.a_type,
		})
	}

	if cmd.Flags().Changed("meta") {
		patches = append(patches, activity.EditField{
			Name:  "meta_label",
			Value: f.meta,
		})
	}

	if cmd.Flags().Changed("date-start") {
		patches = append(patches, activity.EditField{
			Name:  "date_start",
			Value: f.date_start,
		})
	}

	if cmd.Flags().Changed("date-end") {
		patches = append(patches, activity.EditField{
			Name:  "date_end",
			Value: f.date_end,
		})
	}

	return patches
}

func fillArgs(cmd cobra.Command) (cobra.Command, *flagVars) {
	vars := flagVars{}
	cmd.Flags().StringVar(&vars.name, "name", "", "name for activity")
	cmd.Flags().StringVar(&vars.subtitle, "subtitle", "", "subtitle for activity")
	cmd.Flags().StringVar(&vars.description, "description", "", "description for activity")
	cmd.Flags().StringVar(&vars.a_type, "type", "", "type of activity")
	cmd.Flags().StringVar(&vars.meta, "meta", "", "meta label for activity")
	cmd.Flags().StringVar(&vars.date_start, "date-start", "", "start date of activity")
	cmd.Flags().StringVar(&vars.date_end, "date-end", "", "end date of activity")
	return cmd, &vars
}

func ActivityCmd(repo activity.ActivityRepository) *cobra.Command {
	cmd := cobra.Command{
		Use:   "activity <id>",
		Short: "Update activity fields by ID",
		Args:  cobra.ExactArgs(1),
	}
	cmd, vars := fillArgs(cmd)
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return err
		}
		patches := getPatches(cmd, vars)
		_, err = repo.Edit(id, patches)
		return err
	}
	return &cmd
}
