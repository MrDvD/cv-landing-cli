package cmd

import (
	"cv-landing-cli/cmd/add"
	"cv-landing-cli/cmd/edit"
	"cv-landing-cli/cmd/remove"
	"cv-landing-cli/cmd/show"
	"cv-landing-cli/pkg/activity"
	"cv-landing-cli/pkg/attachments"
	"cv-landing-cli/pkg/client"
	"cv-landing-cli/pkg/config"
	"cv-landing-cli/pkg/model/action"
	"cv-landing-cli/pkg/model/history"
	"cv-landing-cli/pkg/model/shell"
	"cv-landing-cli/pkg/tags"
	"net/http"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

var rootCmd = cobra.Command{
	Use:   "cv-admin",
	Short: "admin tool for cv management",
	Long: `The administrative companion to the CV Landing CLI.
		
CV Admin allows you to manage your professional identity directly 
from the terminal, providing interactive forms for data entry and 
automated formatting for your deployment-ready CV.`,
}
var baseClient = client.BaseClient{
	HTTPClient: http.DefaultClient,
	Endpoints:  config.MustGetAppConfig().Hosts,
}
var activityClient = activity.ActivityClient{
	Base: &baseClient,
}
var attachmentClient = attachments.AttachmentClient{
	Base: &baseClient,
}
var tagsClient = tags.TagsClient{
	Base: &baseClient,
}

func fillModes(cmd cobra.Command) (cobra.Command, *bool) {
	interactive := new(bool)
	cmd.Flags().BoolVarP(interactive, "interactive", "i", false, "enable interactive mode")
	return cmd, interactive
}

func Execute() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.AddCommand(add.InitCmd(&activityClient))
	rootCmd.AddCommand(show.InitCmd(&activityClient, &attachmentClient, &tagsClient))
	rootCmd.AddCommand(remove.InitCmd(&activityClient))
	rootCmd.AddCommand(edit.InitCmd(&activityClient))
	rootCmd, interactive := fillModes(rootCmd)
	rootCmd.RunE = func(cmd *cobra.Command, args []string) error {
		if *interactive {
			action := action.NewModel("Action menu", "press q to quit")
			history := history.New()
			shell := shell.NewShell(history.Push(action))
			p := tea.NewProgram(shell)
			_, err := p.Run()
			return err
		} else {
			return rootCmd.Help()
		}
	}
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
