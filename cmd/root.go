package cmd

import (
	"cv-landing-cli/cmd/add"
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

var rootCmd = &cobra.Command{
	RunE: func(cmd *cobra.Command, args []string) error {
		p := tea.NewProgram(action.NewModel())
		_, err := p.Run()
		return err
	},
}
var activityClient = activity.ActivityClient{
	Base: &client.BaseClient{
		HTTPClient: http.DefaultClient,
		Endpoints:  config.MustGetAppConfig().Hosts,
	},
}

func Execute() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.AddCommand(add.InitCmd(&activityClient))
	rootCmd.AddCommand(show.InitCmd(&activityClient))
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
