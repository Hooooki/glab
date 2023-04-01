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

package config

//go:generate bash -c "go install ../tools/yannotated && yannotated -o sample.go -format go -package config -type Config"

import (
	"gopkg.in/yaml.v3"
	"io"
	"os"
	"path/filepath"
)

var (
	// FileName ConfigFileName stores file of config
	FileName = ".config"
)

type Context struct {
	Id             int `yaml:"id"`
	CurrentProject Project
}

type Project struct {
	Id    int    `yaml:"id"`
	Name  string `yaml:"name"`
	Token string `yaml:"token"`
}

type Config struct {
	Projects []Project `yaml:"projects"`
	Context  Context   `yaml:"context"`
	DryRun   bool      `yaml:"dryRun,omitempty"`
}

// New creates new config object
func New() (*Config, error) {
	c := &Config{}
	if err := c.Load(); err != nil {
		return c, err
	}

	return c, nil
}

func createIfNotExist() error {
	// create file if not exist
	configFile := filepath.Join(configDir(), FileName)
	_, err := os.Stat(configFile)
	if err != nil {
		if os.IsNotExist(err) {
			file, err := os.Create(configFile)
			if err != nil {
				return err
			}
			file.Close()
		} else {
			return err
		}
	}

	return nil
}

// Load loads configuration from config file
func (c *Config) Load() error {
	err := createIfNotExist()
	if err != nil {
		return err
	}

	file, err := os.Open(getConfigFile())
	if err != nil {
		return err
	}

	b, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	if len(b) != 0 {
		return yaml.Unmarshal(b, c)
	}

	return nil
}

func (c *Config) Write() error {
	f, err := os.OpenFile(getConfigFile(), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	enc := yaml.NewEncoder(f)
	enc.SetIndent(2) // compat with old versions of kubewatch
	return enc.Encode(c)
}

func getConfigFile() string {
	configFile := filepath.Join(configDir(), FileName)
	if _, err := os.Stat(configFile); err == nil {
		return configFile
	}

	return ""
}

func configDir() string {

	configDir := os.Getenv("GLAB_DIRECTORY")

	if configDir == "" {
		configDir = os.Getenv("HOME") + "/.config/glab"
	}

	if _, err := os.Stat(configDir); os.IsNotExist(err) {
		err := os.MkdirAll(configDir, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}

	return configDir
}
