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
	"fmt"
	"github.com/spf13/cobra"
	"slog"
)


// To be debated > go the cron route or just simply write out from YAML config file


var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create a dump",
	Long: `Creates a dump from a config file. 

All values can be assigned directly into the config file. For templates and a more detailed reference, see the docs`,
	Run: runCreate,
}

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.Flags().StringP("conf", "C", "", "config file with dump parameters (required)")
}

func runSchedule(cmd *cobra.Command, _ []string) {
	f, err := cmd.Flags().GetString("conf")
	if err != nil {
		slog.Error(err, "Command failed to execute")
	}

	slog.Info("The following file will be written: " + f)
	
	
}
