package history

import "cv-landing-cli/pkg/model/shell"

type HistoryModel interface {
	shell.ShellContent
	Push(child shell.ShellContent) HistoryModel
	Current() shell.ShellContent
}
