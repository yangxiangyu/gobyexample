package cmd

import (
	"fmt"

	"time"

	"github.com/spf13/cobra"
)

// limitrateCmd represents the limitrate command
var limitrateCmd = &cobra.Command{
	Use:   "limitrate",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		requests := make(chan int, 5)
		for i := 1; i <= 5; i++ {
			requests <- i
		}
		close(requests)
		rate := time.Second / 10
		limiter := time.Tick(rate)
		fmt.Println("limitrate called")
		for req := range requests {
			<-limiter
			fmt.Println("request", req, time.Now())
		}
	},
}

func init() {
	RootCmd.AddCommand(limitrateCmd)
}
