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

// Package main is the entrypoint for tabled
package main

import (
	"log"
	"strings"

	"github.com/nyrahul/tabled/pkg/config"
	"github.com/nyrahul/tabled/pkg/drawtable"
	"github.com/nyrahul/tabled/pkg/version"
	"github.com/spf13/cobra"
)

var cfg config.Config

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
	rootCmd.PersistentFlags().StringVar(&cfg.Title, "title", "", "title to be used for the table")
	rootCmd.PersistentFlags().StringVar(&cfg.Caption, "caption", "", "caption to be used for the table")
}

func main() {
	log.Printf("version: %s\n", version.Version)
	cobra.CheckErr(rootCmd.Execute())
	if strings.HasSuffix(cfg.InFile, ".csv") {
		drawtable.Csv2Table(cfg)
	} else {
		log.Fatal("--in option not provided or unsupported file extn")
	}
}
