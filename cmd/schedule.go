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
)


// To be debated > go the cron route or just simply write out from YAML config file


var scheduleCmd = &cobra.Command{
	Use:   "schedule",
	Short: "Schedules a dump",
	Long: `Schedules a dump. 

Notice: for this command to work properly, it is required to have set the frequency and path parameters in btool's configuration file, located at $HOME/.config/btool/`,
	Run: runSchedule,
}

func init() {
	rootCmd.AddCommand(scheduleCmd)
	logCmd.Flags().StringP("file", "F", "", "file name to write the dump to")
}

func runSchedule(cmd *cobra.Command, _ []string) {
	f, err := cmd.Flags().GetString("file")
	if err != nil {
		slog.Error(err, "Command failed to execute")
	}

	slog.Info("The following file will be written: " + f)
	
	
}
