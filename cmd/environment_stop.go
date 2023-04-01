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
	"gitlab-environment/pkg/client"
	"strconv"
)

var StopEnvironmentCmd = &cobra.Command{
	Use:       "stop [name]",
	Short:     "Stop an environment",
	Long:      "Stop an environment for the current project",
	ValidArgs: []string{"name"},
	Aliases:   []string{"s"},
	Args:      cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		id, err := strconv.ParseInt(args[0], 10, 0)

		if nil != err {
			color.Red(err.Error())
			return
		}

		e, err := client.StopEnvironment(cfg, int(id))

		if nil != err {
			color.Red(err.Error())
			return
		}

		color.Yellow("Environment %s has been stopped", e.Name)
	},
}
