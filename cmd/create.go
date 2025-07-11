/*
Copyright 2025 The btool Authors All rights reserved

DISCLAIMER:

btool is free software; you can redistribute it and/or modify it under
the terms of the GNU General Public License as published by the Free
Software Foundation

btool is distributed in the hope that it will be useful, but WITHOUT ANY
WARRANTY; without even the implied warranty of MERCHANTABILITY or
FITNESS FOR A PARTICULAR PURPOSE.

The full license can be seen in the file ./LICENSE.  If not see
<http://www.gnu.org/licenses/>.
*/

package cmd

import (
	"github.com/spf13/cobra"
	"log/slog"
	"os"
	"os/exec"
)

// rootCmd represents the base command when called without any subcommands
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates a dump routine",
	Long: `Creates a dump routine. 

Dumps will be created according to the options specified in the configuration file located in $HOME/.config/btool/conf.yaml`,
	Run: runCreate,
}

func init() {

	rootCmd.AddCommand(createCmd)
	createCmd.Flags().StringP("file", "F", "", "file name to write dump to (required)")

	err := createCmd.MarkFlagRequired("file")
	if err != nil {
		slog.Error(err.Error())
	}

}

func runCreate(cmd *cobra.Command, args []string) {

	// Change directory
	uHome := os.Getenv("HOME")
	p := uHome
	// Get the directory to save from conf.yaml in the Dumps.path section

	err := os.Chdir(p)
	if err != nil {
		slog.Error(err.Error())
	}

	// *** The following block is meant only for testing ***

	wDir, _ := os.Getwd()
	if err != nil {
		slog.Error("Couldn't change working directory: " + err.Error())
	}
	slog.Info("Setting Working directory as: " + wDir)

	cmdStr := "mariadb-dump mariadb_test > test-dump1.sql"
	slog.Info("Executing: " + cmdStr)
	c := exec.Command("/bin/sh", "-c", "mariadb-dump mariadb_test > test-dump1.sql")

	var stdout []byte
	stdout, err = c.Output()

	if err != nil {
		slog.Error(err.Error())
	}

	// Print the output
	slog.Info(string(stdout))

	// *** End of Testing ***

}
