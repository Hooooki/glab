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
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
	"gitlab-environment/pkg/utils"
)

var showToken bool

var ContextListCmd = &cobra.Command{
	Use:     "list",
	Short:   "List glab contexts",
	Long:    "List glab contexts",
	Aliases: []string{"ls"},
	Run: func(cmd *cobra.Command, args []string) {

		t := *utils.Table(table.Row{"", "Id", "Name", "Token"})

		for _, p := range cfg.Projects {

			token := fmt.Sprintf("%s****************", string([]rune(p.Token)[:len(p.Token)-16]))

			if showToken {
				token = p.Token
			}

			isCurrent := ""

			if cfg.Context.Id == p.Id {
				isCurrent = "*"
			}

			t.AppendRows([]table.Row{
				{isCurrent, p.Id, p.Name, token},
			})

		}

		t.Render()
	},
}

func init() {
	ContextListCmd.Flags().BoolVar(&showToken, "show-tokens", false, "Show project's token")
}
