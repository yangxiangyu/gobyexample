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
	"time"
	"strconv"
)

// timeCmd represents the time command
var timeCmd = &cobra.Command{
	Use:   "time",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		p := fmt.Println

		now := time.Now()
		p(now)

		d := time.Duration(7200 * 1000 * 1000 * 1000)
		p(d)

		then := time.Date(
			2013, 1, 7, 20, 34, 58, 651387237, time.UTC)

		p(then)
		p(then.Year())
		p(then.Month())
		p(then.Day())
		p(then.Hour())
		p(then.Minute())
		p(then.Second())
		p(then.Nanosecond())
		p(then.Location())
		p(then.Weekday())

		p(then.Before(now))
		p(then.After(now))
		p(then.Equal(now))

		p(then.Date())
		p(then.ISOWeek())
		p("----------")
		p(now.UTC())
		p(now.Local())
		p(now.Location())
		p(now.Zone())
		p(now.Unix())
		p(time.Unix(now.Unix(), 0))
		p(now.UnixNano())
		p(time.Unix(0, now.UnixNano()))
		p(now.GobEncode())
		p(now.MarshalJSON())
		p(time.Since(now))
		p("----------")
		diff := now.Sub(then)
		p(diff)

		p(diff.Hours())
		p(diff.Minutes())
		p(diff.Seconds())
		p(diff.Nanoseconds())
		p(then.Add(diff))
		p(then.Add(-diff))

		p(d)
		p(d.Hours())
		p(d.Minutes())
		p(d.Seconds())
		p(d.Nanoseconds())
		p(then.Add(d))
		fmt.Println("time now")
		fmt.Println(strconv.FormatInt(time.Now().Unix(), 10))
		fmt.Println("time called")
	},
}

func init() {
	RootCmd.AddCommand(timeCmd)

}
