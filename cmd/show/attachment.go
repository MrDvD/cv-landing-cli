package show

import (
	"cv-landing-cli/pkg/attachments"
	"cv-landing-cli/pkg/view"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var attachmentTableConfig = view.TableConfig[attachments.Attachment]{
	Columns: []view.Column{
		{Header: "ID", Weight: 1},
		{Header: "NAME", Weight: 3},
		{Header: "LINK", Weight: 3},
		{Header: "PRIORITY", Weight: 1},
		{Header: "ACTIVITY_ID", Weight: 1},
	},
	RowMapper: func(a attachments.Attachment) []string {
		return []string{
			fmt.Sprintf("%d", a.Id),
			a.Name,
			a.Link,
			view.ShowPtr(a.Priority),
			fmt.Sprintf("%d", a.ActivityId),
		}
	},
}

func AttachmentCmd(repo attachments.AttachmentRepository) *cobra.Command {
	cmd := cobra.Command{
		Use:  "attachment <id>",
		Args: cobra.ExactArgs(1),
	}
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return err
		}
		attachments, err := repo.Get(id)
		if err != nil {
			return err
		}
		view.Table(attachments, attachmentTableConfig)
		return nil
	}
	return &cmd
}
