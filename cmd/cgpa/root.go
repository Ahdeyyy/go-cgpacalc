package cgpa

import (
	"cgpa-calc/database"
	"cgpa-calc/ui"
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

var db database.CgpaRepo = database.NewCgpaRepo("cgpa.db")

var rootCmd = &cobra.Command{
	Use:   "cgpa",
	Short: "cgpa - calculating your cgpa",
	Long:  "cgpa - calculating your cgpa",
	Run: func(cmd *cobra.Command, args []string) {
		sessions, _ := db.GetSemesters()
		items := make([]list.Item, len(sessions))
		for i, session := range sessions {
			items[i] = session
		}

		m := ui.ListSessionModel{
			List: list.New(items, list.NewDefaultDelegate(), 0, 0),
		}
		m.List.Title = "My Fave Things"

		p := tea.NewProgram(m, tea.WithAltScreen())

		if _, err := p.Run(); err != nil {
			fmt.Println("Error running program:", err)
			os.Exit(1)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}
