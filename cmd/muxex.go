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
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"net/http"
)

// muxexCmd represents the muxex command
var muxexCmd = &cobra.Command{
	Use:   "muxex",
	Short: "mux 自定义路由示例",
	Run: func(cmd *cobra.Command, args []string) {
		run()
		fmt.Println("muxex called")
	},
}

func init() {
	RootCmd.AddCommand(muxexCmd)
}

func run() {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w,"hahaha")
	})
	n := negroni.Classic()
	n.UseHandler(r)
	n.Run(":3000")
}
