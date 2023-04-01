/* MIT License

Copyright (c) 2023 Fragan Gourvil

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE */

package cmd

import (
	"github.com/fatih/color"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
	"gitlab-environment/pkg/client"
	"gitlab-environment/pkg/utils"
)

var ListEnvironmentCmd = &cobra.Command{
	Use:     "list",
	Short:   "List environments",
	Long:    "List each environment for the current project",
	Aliases: []string{"ls"},
	Run: func(cmd *cobra.Command, args []string) {

		environments, err := client.ListEnvironment(cfg)

		if nil != err {
			color.Red(err.Error())
			return
		}

		t := *utils.Table(table.Row{"#", "Name", "Slug", "State", "URL", "Tier", "Created at", "Updated at"})

		for _, e := range *environments {

			t.AppendRows([]table.Row{
				{e.Id, e.Name, e.Slug, e.State, e.ExternalUrl, e.Tier, e.CreatedAt, e.UpdatedAt},
			})

		}

		t.Render()

	},
}

func init() {
}
