// Copyright © 2017 Jimmy Song
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
	"github.com/rootsongjc/magpie/utils"
	"github.com/spf13/cobra"
)

var host string
var detail bool

// lookupCmd represents the lookup command
var lookupCmd = &cobra.Command{
	Use:   "lookup",
	Short: "Lookup a nodemanager container",
	Long:  "Look up the nodemanager docker container.",
	Run: func(cmd *cobra.Command, args []string) {
		utils.Lookup(host, detail)
	},
}

func init() {
	toolCmd.AddCommand(lookupCmd)

	lookupCmd.Flags().StringVarP(&host, "hostname", "i", "", "Look up the nodemanager")
	lookupCmd.Flags().BoolVarP(&detail, "detail", "d", false, "Show nodemanager container's detail messages")

}
