package cmd

import (
	"cv-landing-cli/cmd/add"
	"cv-landing-cli/pkg/activity"
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
	Client:  http.DefaultClient,
	ApiLink: config.MustGetAppConfig().ActivityApiBase,
}

func Execute() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.AddCommand(add.InitCmd(&activityClient))
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
