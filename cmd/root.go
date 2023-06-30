/*
Copyright 2023 Rahul Jadhav

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Package cmd is the entrypoint for tabled
package cmd

import (
	"log"
	"strings"

	"github.com/nyrahul/tabled/pkg/config"
	"github.com/nyrahul/tabled/pkg/drawtable"
	"github.com/spf13/cobra"
)

var cfg config.Config
var yamlFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Run: func(cmd *cobra.Command, args []string) {
	},
	Use:           "tabled",
	Short:         "Table Designer",
	Long:          `CLI Utility to help plot a table from csv/json input`,
	SilenceUsage:  true,
	SilenceErrors: true,
}

func init() {
	rootCmd.PersistentFlags().StringVar(&cfg.InFile, "in", "", "file whose contents are to be plotted in the table")
	rootCmd.PersistentFlags().StringVar(&yamlFile, "config", "", "configuration to be used (yaml)")
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
	cfg.YamlCfg = config.LoadYAMLConfig(yamlFile)
	if strings.HasSuffix(cfg.InFile, ".csv") {
		drawtable.Csv2Table(cfg)
	} else {
		log.Fatal("--in option not provided or unsupported file extn")
	}
}
