package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

type tesstruct struct {
	p string
}

// blackboardCmd represents the blackboard command
var blackboardCmd = &cobra.Command{
	Use:   "blackboard",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		f1()
		fmt.Println("blackboard called")
	},
}

func init() {
	RootCmd.AddCommand(blackboardCmd)

}

func f1() {
	ts := tesstruct{p: "ssss"}
	fmt.Println("aaaa")
	fmt.Println(ts)

}
