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
	"github.com/spf13/cobra"
	"gitlab-environment/pkg/config"
	"strconv"
)

var ContextAddCmd = &cobra.Command{
	Use:       "add [name] [id] [token]",
	Short:     "Add a project to glab contexts",
	Long:      "Add a project to glab contexts",
	ValidArgs: []string{"name", "id", "token"},
	Aliases:   []string{"a"},
	Args:      cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {

		id, err := strconv.ParseInt(args[1], 10, 0)

		if nil != err {
			color.Red(err.Error())
			return
		}

		p := config.Project{
			Id:    int(id),
			Name:  args[0],
			Token: args[2],
		}

		cfg.Projects = append(cfg.Projects, p)
		err = cfg.Write()

		if nil != err {
			color.Red(err.Error())
			return
		}

		color.Green("Context %d (%s) has been added to the configuration", p.Id, p.Name)
	},
}
