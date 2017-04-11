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
	"fmt"

	"github.com/spf13/cobra"
)

// hashsetCmd represents the hashset command
var hashsetCmd = &cobra.Command{
	Use:   "hashset",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		s := New()
		s.Add(1)
		s.Add(1)
		s.Add(1)
		b := s.Contains(1)
		fmt.Println(b)
		fmt.Println("hashset called")
	},
}

func init() {
	RootCmd.AddCommand(hashsetCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// hashsetCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// hashsetCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}

var itemExists = struct{}{}

type Set struct {
	items map[interface{}]struct{}
}

func New() *Set {
	return &Set{items: make(map[interface{}]struct{})}
}

func (set *Set) Add(item interface{}) {
	set.items[item] = itemExists
}

func (set *Set) Remove(item interface{}) {
	delete(set.items, item)
}

func (set *Set) Contains(item interface{}) bool {
	if _, contains := set.items[item]; !contains {
		return false
	}
	return true
}
