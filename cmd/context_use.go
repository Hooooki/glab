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
	"errors"
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"strconv"
)

var ContextUseCmd = &cobra.Command{
	Use:       "use [id|name]",
	Short:     "Switch context",
	Long:      "Switch your context to another project",
	ValidArgs: []string{"name", "id"},
	Aliases:   []string{"u"},
	Args:      cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		id, err := strconv.ParseInt(args[0], 10, 0)

		if nil != err {

			if nil != switchWithName(args[0]) {
				color.Red(err.Error())
				return
			}

		}

		if nil != switchWithId(int(id)) {
			if err := cfg.Write(); err != nil {
				color.Red(err.Error())
				return
			}
		}

		err = cfg.Write()
		if err != nil {
			color.Red(err.Error())
		}

	},
}

func switchWithId(id int) error {

	found := false

	for _, p := range cfg.Projects {

		if p.Id == id {
			cfg.Context.CurrentProject = p
			cfg.Context.Id = p.Id
			found = true
		}

	}

	if !found {
		return errors.New(fmt.Sprintf("Context %d doesn't exist", id))
	}

	fmt.Printf("Switched to %s", cfg.Context.CurrentProject.Name)

	return nil
}

func switchWithName(name string) error {

	found := false

	for _, p := range cfg.Projects {

		if p.Name == name {
			cfg.Context.CurrentProject = p
			cfg.Context.Id = p.Id
			found = true
		}

	}

	if !found {
		return errors.New(fmt.Sprintf("Context %s doesn't exist", name))
	}

	fmt.Printf("Switched to %s", cfg.Context.CurrentProject.Name)

	return nil
}
