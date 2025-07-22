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

/*
   IMPORTANT COMMENT: 
   
   With the use of parsing config files in general, there is no use anymore for this command,
   however we will substitute this for a config Template creation method, such as:
   
   btool-config init
   
   This command will create a configuration file in the current directory

*/


package cmd

import (
	"log/slog"
	"os"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializes a sample configuration YAML file in the current directory",
	Long: `Initializes a sample configuration YAML file in the current directory, parsing the "name" flag as the file name (required). 
	
Exampĺe usage: btool init --file="btool_config"
	
This will create an empty config file btool_config.yaml in the current directory, which is used for later creating dumps/backups.
	
This file contains all the available configuration options you could set values to. For more information on how to make the most out of the file type, see docs.`,
	Run: runInit,
}

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().StringP("file", "f", "", "name of the file to be created (required)")
	
	err := initCmd.MarkFlagRequired("file")
	if err != nil {
		slog.Error(err.Error())
	}
}

func runInit(cmd *cobra.Command, args []string) {
	
	f, err := cmd.Flags().GetString("file")
	if err != nil {
		slog.Error("File %s already exists", f)
	}
	
	if cFile(f) == false {
		slog.Info("File does not exist, creating...")

		var content =
		// Template file
		"# BTOOL - A general-purpose backup tool to create dumps via configuration files\n" +
"# -----------------------------------------------------------------------------------------\n" +
"# This is the configuration template file, it contains all possible attributes that btool\n" +
"# can pull information from to properly process backup requests. For a detailed description\n" +
"# of how to use each attribute, please see the spec file in the documentation at: \n" +
"# https://github.com/pages/etc/etc\n" +
"database:\n" +
" engine:\n" +
" user:\n" +
" password:\n" + 
" host:\n" + 
"dump:\n" +
" path:\n" +
" type:\n" +
" database_name:\n" + 
" tables:\n"
    

		// Creating the configuration file
		err := os.WriteFile(f, []byte(content), 0755)
		if err != nil {
			slog.Error(err.Error())
		}

	} else {
		slog.Info("Configuration file created!")
	}

	slog.Info("Success - Init completed. Run 'btool --help' for usage.")
	
}

func cFile(filename string) bool {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return true
}


