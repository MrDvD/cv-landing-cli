package cmd

import (
	"cv-landing-cli/cmd/add"
	"cv-landing-cli/cmd/edit"
	"cv-landing-cli/cmd/remove"
	"cv-landing-cli/cmd/show"
	"cv-landing-cli/pkg/activity"
	"cv-landing-cli/pkg/client"
	"cv-landing-cli/pkg/config"
	"cv-landing-cli/pkg/model/action"
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
var activityClient = activity.ActivityClient{
	Base: &client.BaseClient{
		HTTPClient: http.DefaultClient,
		Endpoints:  config.MustGetAppConfig().Hosts,
	},
}

func fillModes(cmd cobra.Command) (cobra.Command, *bool) {
	interactive := new(bool)
	cmd.Flags().BoolVarP(interactive, "interactive", "i", false, "enable interactive mode")
	return cmd, interactive
}

func Execute() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.AddCommand(add.InitCmd(&activityClient))
	rootCmd.AddCommand(show.InitCmd(&activityClient))
	rootCmd.AddCommand(remove.InitCmd(&activityClient))
	rootCmd.AddCommand(edit.InitCmd(&activityClient))
	rootCmd, interactive := fillModes(rootCmd)
	rootCmd.RunE = func(cmd *cobra.Command, args []string) error {
		if *interactive {
			p := tea.NewProgram(action.NewModel())
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
