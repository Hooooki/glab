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
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	c "gitlab-environment/pkg/config"
	"os"
)

var projectId int
var privateToken string
var dryRun bool
var cfgFile string

var cfg = &c.Config{}

var RootCmd = &cobra.Command{
	Use:   "glab",
	Short: "Manage your Gitlab projects",
	Long:  "GLab allow you to manage your Gitlab projects easily",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if err := cfg.Load(); err != nil {
			panic(err)
		}

		for _, p := range cfg.Projects {

			if cfg.Context.Id == p.Id {
				cfg.Context.CurrentProject = p
			}

		}

		if 0 != projectId {
			cfg.Context.CurrentProject.Id = projectId
		}

		if "" != privateToken {
			cfg.Context.CurrentProject.Token = privateToken
		}

	},
}

func init() {

	RootCmd.SetHelpCommand(&cobra.Command{
		Use:    "no-help",
		Hidden: true,
	})

	cobra.OnInitialize(initConfig)

	RootCmd.PersistentFlags().IntVarP(&projectId, "project-id", "p", 0, "Gitlab project ID (required)")
	RootCmd.PersistentFlags().StringVarP(&privateToken, "private-token", "t", "", "Gitlab user private token (required)")
	RootCmd.PersistentFlags().BoolVarP(&dryRun, "dry-run", "d", false, "Run command in debug mode.")

	RootCmd.AddCommand(EnvironmentCmd, ContextCmd)
}

// initConfig reads in cfg file and ENV variables if set.
func initConfig() {
	if cfgFile != "" { // enable ability to specify cfg file via flag
		viper.SetConfigFile(cfgFile)
	}

	viper.SetConfigName(".glab") // name of cfg file (without extension)
	viper.AddConfigPath("$HOME") // adding home directory as first search path
	viper.AutomaticEnv()         // read in environment variables that match

	// If a cfg file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using cfg file:", viper.ConfigFileUsed())
	}
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
