package show

import (
	"cv-landing-cli/pkg/tags"
	"cv-landing-cli/pkg/view"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

type flagArgs struct {
	tagId   string
	tagType string
}

var tagsTableConfig = view.TableConfig[tags.Tag]{
	Columns: []view.Column{
		{Header: "ID", Weight: 1},
		{Header: "NAME", Weight: 3},
		{Header: "TYPE", Weight: 2},
		{Header: "PRIORITY", Weight: 1},
		{Header: "ACTIVITY_ID", Weight: 1},
	},
	RowMapper: func(t tags.Tag) []string {
		return []string{
			fmt.Sprintf("%d", t.Id),
			t.Name,
			t.Type,
			view.ShowPtr(t.Priority),
			fmt.Sprintf("%d", t.ActivityId),
		}
	},
}

func getFilter(cmd *cobra.Command, f *flagArgs) (tags.TagFilter, error) {
	filter := tags.TagFilter{
		TagType: &f.tagType,
	}

	if cmd.Flags().Changed("activity-id") {
		id, err := strconv.Atoi(f.tagId)
		if err != nil {
			return tags.TagFilter{}, err
		}
		filter.ActivityID = &id
	}

	return filter, nil
}

func fillVars(cmd cobra.Command) (cobra.Command, *flagArgs) {
	var args flagArgs
	cmd.Flags().StringVar(&args.tagType, "type", "", "type of tags")
	cmd.MarkFlagRequired("type")
	cmd.Flags().StringVar(&args.tagId, "activity-id", "", "type of tags")
	return cmd, &args
}

func TagCmd(repo tags.TagsRepository) *cobra.Command {
	cmd := cobra.Command{
		Use: "tag",
	}
	cmd, vars := fillVars(cmd)
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		filter, err := getFilter(cmd, vars)
		if err != nil {
			return err
		}
		tags, err := repo.Get(filter)
		if err != nil {
			return err
		}
		view.Table(tags, tagsTableConfig)
		return nil
	}
	return &cmd
}
