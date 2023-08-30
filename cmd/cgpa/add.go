package cgpa

import (
	"cgpa-calc/ui"
	"fmt"
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

var isSemester bool = false
var isCourse bool = false

var addSemesterCmd = &cobra.Command{
	Use:     "add",
	Aliases: []string{"a"},
	Short:   "adds a new semester or course to the database by using the --semester or --course flag.",
	Long:    "adds a new semester or course to the database by using the --semester or --course flag, if the course and semester flags are set semester would be priotized above course.",
	Run: func(cmd *cobra.Command, args []string) {
		
		if isSemester {
			p := tea.NewProgram(ui.InitialCSModel(db))
			if _, err := p.Run(); err != nil {
				log.Fatal(err)
			}
		}
		
		if isCourse {
			fmt.Println("adding a course is comming soon...")
		}

		if !isSemester && !isCourse {
			fmt.Println("please specify what you want to add, either a semester or a course with the --semester or --course flag")
		}
	},
}

func init() {
	rootCmd.AddCommand(addSemesterCmd)
	addSemesterCmd.PersistentFlags().BoolVarP(&isSemester,"semester", "s", false, "Create and add a new semester to the database")
	
	addSemesterCmd.PersistentFlags().BoolVarP(&isCourse,"course","c", false, "Create and add a new course to the database")

}
