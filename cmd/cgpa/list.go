package cgpa

import (
	"cgpa-calc/ui"
	"fmt"

	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"l"},
	Short:   "lists the semesters or courses in the database by using the --semester or --course flag.",
	Long:    "lists the semesters or courses in the database by using the --semester or --course flag, if the course and semester flags are set semester would be priotized above course.",
	Run: func(cmd *cobra.Command, args []string) {

		if isSemester {
			sessions, _ := db.GetSemesters()
			items := make([]list.Item, len(sessions))
			for i, session := range sessions {
				items[i] = session
			}

			m := ui.ListSessionModel{
				List: list.New(items, list.NewDefaultDelegate(), 0, 0),
			}
			m.List.Title = "Your Semesters"

			p := tea.NewProgram(m, tea.WithAltScreen())

			if _, err := p.Run(); err != nil {
				fmt.Println("Error running program:", err)
				os.Exit(1)
			}
		}

		if isCourse {
			courses, _ := db.GetAllCourses()
			items := make([]list.Item, len(courses))
			for i, course := range courses {
				items[i] = course
			}

			m := ui.ListSessionModel{
				List: list.New(items, list.NewDefaultDelegate(), 0, 0),
			}

			m.List.Title = "Your Courses"

			p := tea.NewProgram(m, tea.WithAltScreen())

			if _, err := p.Run(); err != nil {
				fmt.Println("Error running program:", err)
				os.Exit(1)
			}

		}

		if !isSemester && !isCourse {
			fmt.Println("please specify what you want to listed, either a semester or a course with the --semester or --course flag")
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.PersistentFlags().BoolVarP(&isSemester, "semester", "s", false, "Create and add a new semester to the database")

	listCmd.PersistentFlags().BoolVarP(&isCourse, "course", "c", false, "Create and add a new course to the database")

}
