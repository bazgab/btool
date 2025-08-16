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
	"gopkg.in/yaml.v3"
	"log/slog"
	"os/exec"
	"os"
	"strings"
)


var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create a dump",
	Long: `Creates a dump from a config file. 

All values can be assigned directly into the config file. For templates and a more detailed reference spec sheet, see: 
https://github.com/bazgab/btool/etc/etc`,
	Run: runCreate,
}

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.Flags().StringP("config", "C", "", "config file with dump parameters (required)")
}

func runCreate(cmd *cobra.Command, _ []string) {
	c, err := cmd.Flags().GetString("config")
	if err != nil {
		slog.Error(err.Error())
	}

	slog.Info("The following file has been selected for config values: " + c)

	
	type Config struct {
		Database struct {
			Engine   string `yaml:"engine"`
			User     string `yaml:"user"`
			Password string `yaml:"password"`
			Host     string `yaml:"host"`
		} `yaml:"database"`
		Dump struct {
			Path         string   `yaml:"path"`
			Type         string   `yaml:"type"`
			DatabaseName []string `yaml:"database_name"`
			Tables       []string `yaml:"tables"`
		} `yaml:"dump"`
	}

	confFile, err := os.ReadFile(c)
	if err != nil {
		fmt.Println(err.Error())
	}

	replaced := os.ExpandEnv(string(confFile))

	var confValues Config
	err = yaml.Unmarshal([]byte(replaced), &confValues)
	if err != nil {
		fmt.Println(err.Error())
	}

	/* Testing for multiple values 
	for i := 0; i < len(confValues.Dump.DatabaseName); i++ {
		fmt.Printf("Database name %d : %s\n", i, confValues.Dump.DatabaseName[i])

	}
	*/
	
	// Design Issue:
	// overcomplicating by having default values. 
	
	// An approach to not have multiple values laying around:
	
	// Perform a basic dump
	
	for i := 0; i < len(confValues.Dump.DatabaseName); i++ {
		fmt.Printf("Creating dump for : %s\n", confValues.Dump.DatabaseName[i])
		dumpName := confValues.Dump.DatabaseName[i] + "-dump.sql"
		arg := []string{"--databases", confValues.Dump.DatabaseName[i], ">", dumpName }
		fmt.Sprintf("Running command with following argument: %s", strings.Join(arg, " "))
		out, err := exec.Command("usr/bin/mariadb-dump", strings.Join(arg, " ")).Output()
		if err != nil {
			slog.Info("Error when executing dump command")
			os.Exit(1)
		} 
		fmt.Printf("Output: %s", out)
	}
	
	
	/*
	err := exec.Command("usr/bin/mariadb-dump", arg1).Run()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Command successfully executed")
	}

	*/
	
	/*
	// Firstly check the correct usage of db engine
	if confValues.Database.Engine == "mariadb" {
		slog.Info("MariaDB selected - proceeding setup")
	} else {
		slog.Error("Unsupported Database Engine")
		os.Exit(1)
	}

	//Check values, if not > parse defaults
	if confValues.Database.User == "" {
		slog.Info("User check - No selected user. Setting option to default value: ", defaultUser)
	}	
	slog.Info(fmt.Sprintf("User: %s\n", confValues.Database.User))
	
	if confValues.Database.Password == "" {
		slog.Info("Password check - No password entered. Password is a \033[1mrequired field\033[0m, for usage see the docs.")
		// TODO > if no password selected, panic
	} 
	
	if confValues.Database.Host == "" {
		slog.Info(fmt.Sprintf("Host check - No selected host. Setting option to default value: %s", defaultHost))
	}	
	
	if confValues.Dump.Path == "" {
		slog.Info("Path check - No path entered. Path is a \033[1mrequired field\033[0m, for usage see the docs.")
		// TODO > if no path selected, panic
	} 
	slog.Info(fmt.Sprintf("Creating dump on path: %s", confValues.Dump.Path))
	
	slog.Info("Type check: ")
	
	// A better aproach
	* 
	
	if confValues.Dump.Type != all_databases {
		Enter condition
		
	}
	
	switch confValues.Dump.Type {
		case "all_databases":
			slog.Info("all databases")
		case "select_databases":
			slog.Info("select databases")
		default:
			slog.Info("Invalid/No option selected - using default option: all_databases")
	}
	
	if confValues.Dump.Path == "" {
		slog.Info(fmt.Sprintf("Tables check - No tables entered. Setting option to default value: %s", defaultTables))
		// TODO > if no path selected, panic
	} 

	// Perform dump
	// Added on top for testing functionality 
	
	*/

	
	
}
