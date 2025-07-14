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
	"log/slog"
	"os"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializes btool",
	Long: `Initializes btool. This command is ideally meant to be used only once after installation, 
and provides the instructions for btool to create the necessary directories and files to make
it work as intended.`,
	Run: runInit,
}

func init() {
	rootCmd.AddCommand(initCmd)
}

func runInit(cmd *cobra.Command, args []string) {

	var (
		h              = os.Getenv("HOME")
		btoolDirectory = h + "/.config/btool"
	)

	slog.Info("Checking for " + btoolDirectory + " directory...")

	if _, err := os.Stat(btoolDirectory); os.IsNotExist(err) {
		slog.Warn(btoolDirectory + " directory does not exist")
		slog.Info("Attempting to create it...")

		err := os.MkdirAll(btoolDirectory, 0755)
		if err != nil {
			slog.Error(err.Error())
		}
	} else {
		slog.Info(btoolDirectory + " directory exists")
	}

	f := btoolDirectory + "/conf.yaml"

	slog.Info("Checking for configuration file...")

	if cFile(f) == false {
		slog.Warn("Configuration file does not exist")
		slog.Warn(f + " is required, attempting to create it...")

		var content =
		// Template file
		"---\n" +
			"# Specify database environment options\n" +
			"\n" +
			"Database:\n" +
			"  engine:\n" +
			"  # Please note it is not considered secure to store plain-text values for database credentials.\n" +
			"  # See: https://github.com/bazgab/btool/blob/master/README.md\n" +
			"  user:\n" +
			"  password:\n" +
			"\n" +
			"# Specify dump options\n" +
			"Dumps:\n" +
			"  # Where btool will save the dumps to\n" +
			"  path:\n" +
			"  # How frequently btool will backup the dumps (accepted values are 'hourly', 'daily, or 'weekly')\n" +
			"  frequency:\n"

		// Creating the configuration file
		err := os.WriteFile(f, []byte(content), 0755)
		if err != nil {
			slog.Error(err.Error())
		}

	} else {
		slog.Info("Configuration file exists")
	}

	slog.Info("Success - Init check completed. Run 'btool --help' for usage.")
	
}

func cFile(filename string) bool {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return true
}


