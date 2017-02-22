// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"container/list"
	"fmt"

	"github.com/spf13/cobra"
)

type T struct {
	Id   int64
	Name string
}

// forCmd represents the for command
var forCmd = &cobra.Command{
	Use:   "for",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:


Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		l := list.New()
		l.PushBack("1")
		l.PushBack("2")
		for e := l.Front(); e != nil; e = e.Next() {
			fmt.Println(e.Value)
		}

		data := []int{1, 2, 3, 4, 5}
		ts := []T{{Id: 111, Name: "aaa"}, {Id: 222, Name: "bbbb"}}
		for d := range (data) {
			fmt.Println(d)
		}
		for t:=range(ts){
			fmt.Println(ts[t].Id)
		}

		fmt.Println("for called")
	},
}

func init() {
	RootCmd.AddCommand(forCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// forCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// forCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
