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
		resultc := make(chan int, 10)
		// resp, err := http.Get("http://www.qq.com/")
		// if err != nil {
		// fmt.Println(err)
		// }
		// body, err := ioutil.ReadAll(resp.Body)
		// resultc <- string(body)
		// fmt.Println(i)
		// defer resp.Body.Close()

		go productorbody(resultc)
		go readbody(resultc)
		// close(resultc)
		select {}
		fmt.Println("blackboard called")
	},
}

type UrlList struct {
	urls []string
}

func productorbody(resultc chan<- int) {
	for i := 1; i < 1000; i++ {
		resultc <- i
	}
}

func readbody(resultc chan int) {
	for result := range resultc {
		fmt.Println(result)
	}
}

func init() {
	RootCmd.AddCommand(blackboardCmd)

}
